// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"unicode"

	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ethereum/go-ethereum/log"
	"github.com/naoina/toml"
	"github.com/urfave/cli"
)

const DefaultConfigPath = "./config.toml"
const DefaultKeystorePath = "./keys"

type Config struct {
	Chains       []RawChainConfig `toml:"chains"`
	keystorePath string
}

// RawChainConfig is parsed directly from the config file and should be using to construct the core.ChainConfig
type RawChainConfig struct {
	Name     string            `toml:"name"`
	Type     string            `toml:"type"`
	Id       msg.ChainId       `toml:"id"`       // ChainID
	Endpoint string            `toml:"endpoint"` // url for rpc endpoint
	From     string            `toml:"from"`     // address of key to use
	Opts     map[string]string `toml:"opts"`
}

func NewConfig() *Config {
	return &Config{
		Chains: []RawChainConfig{},
	}
}

// ToTOML writes config to file
func (c *Config) ToTOML(file string) *os.File {
	var (
		newFile *os.File
		err     error
	)

	var raw []byte
	if raw, err = toml.Marshal(*c); err != nil {
		log.Warn("error marshalling toml", "err", err)
		os.Exit(1)
	}

	newFile, err = os.Create(file)
	if err != nil {
		log.Warn("error creating config file", "err", err)
	}
	_, err = newFile.Write(raw)
	if err != nil {
		log.Warn("error writing to config file", "err", err)
	}

	if err := newFile.Close(); err != nil {
		log.Warn("error closing file", "err", err)
	}
	return newFile
}

func (c *Config) validate() error {
	for _, chain := range c.Chains {
		if chain.Type == "" {
			return fmt.Errorf("required field chain.Type empty for chain %d", chain.Id)
		}

		if chain.Endpoint == "" {
			return fmt.Errorf("required field chain.Endpoint empty for chain %d", chain.Id)
		}

		if chain.Name == "" {
			return fmt.Errorf("required field chain.Name empty for chain %d", chain.Id)
		}
	}
	return nil
}

func getConfig(ctx *cli.Context) (*Config, error) {
	var fig Config
	path := DefaultConfigPath
	if file := ctx.GlobalString(ConfigFileFlag.Name); file != "" {
		path = file
	}
	err := loadConfig(path, &fig)
	if err != nil {
		log.Warn("err loading toml file", "err", err.Error())
		return &fig, err
	}
	if ksPath := ctx.GlobalString(KeystorePathFlag.Name); ksPath != "" {
		fig.keystorePath = ksPath
	}
	log.Debug("Loaded config", "path", path)
	err = fig.validate()
	if err != nil {
		return nil, err
	}
	return &fig, nil
}

func loadConfig(file string, config *Config) error {
	fp, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	log.Debug("Loading configuration", "path", filepath.Clean(fp))
	f, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return err
	}
	if err = tomlSettings.NewDecoder(f).Decode(&config); err != nil {
		return err
	}
	return nil
}

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}
