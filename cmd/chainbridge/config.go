package main

import (
	"ChainBridgeV2/core"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"unicode"

	"github.com/ethereum/go-ethereum/log"
	"github.com/naoina/toml"
	"github.com/urfave/cli"
)

const DefaultConfigPath = "./config.toml"
const DefaultKeystorePath = "./keys"

type Config struct {
	Ethereum core.ChainConfig `toml:"ethereum"`
}

func NewConfig() *Config {
	return &Config{
		Ethereum: core.ChainConfig{},
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
