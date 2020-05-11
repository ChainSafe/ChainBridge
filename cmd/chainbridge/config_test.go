// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

func createTempConfigFile() (*os.File, *Config) {
	testConfig := NewConfig()
	ethCfg := RawChainConfig{
		Name:     "chain",
		Type:     "ethereum",
		Id:       "1",
		Endpoint: "endpoint",
		From:     "0x0",
		Opts:     nil,
	}
	testConfig.Chains = []RawChainConfig{ethCfg}
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		fmt.Println("Cannot create temporary file", "err", err)
		os.Exit(1)
	}

	f := testConfig.ToTOML(tmpFile.Name())
	return f, testConfig
}

// Creates a cli context for a test given a set of flags and values
func createCliContext(description string, flags []string, values []interface{}) (*cli.Context, error) {
	set := flag.NewFlagSet(description, 0)
	for i := range values {
		switch v := values[i].(type) {
		case bool:
			set.Bool(flags[i], v, "")
		case string:
			set.String(flags[i], v, "")
		case uint:
			set.Uint(flags[i], v, "")
		default:
			return nil, fmt.Errorf("unexpected cli value type: %T", values[i])
		}
	}
	context := cli.NewContext(nil, set, nil)
	return context, nil
}

func TestLoadConfig(t *testing.T) {
	file, cfg := createTempConfigFile()
	ctx, err := createCliContext("", []string{"config"}, []interface{}{file.Name()})
	if err != nil {
		t.Fatal(err)
	}

	res, err := getConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res.Chains[0], cfg.Chains[0]) {
		t.Errorf("did not match\ngot: %+v\nexpected: %+v", res.Chains[0], cfg.Chains[0])
	}
}

func TestValdiateConfig(t *testing.T) {
	valid := RawChainConfig{
		Name:     "chain",
		Type:     "ethereum",
		Id:       "1",
		Endpoint: "endpoint",
		From:     "0x0",
		Opts:     nil,
	}

	missingType := RawChainConfig{
		Name:     "chain",
		Type:     "",
		Id:       "1",
		Endpoint: "endpoint",
		From:     "0x0",
		Opts:     nil,
	}

	missingEndpoint := RawChainConfig{
		Name:     "chain",
		Type:     "ethereum",
		Id:       "1",
		Endpoint: "",
		From:     "0x0",
		Opts:     nil,
	}

	missingName := RawChainConfig{
		Name:     "",
		Type:     "ethereum",
		Id:       "1",
		Endpoint: "endpoint",
		From:     "0x0",
		Opts:     nil,
	}

	cfg := Config{
		Chains:       []RawChainConfig{valid},
		keystorePath: "",
	}

	err := cfg.validate()
	if err != nil {
		t.Fatal(err)
	}

	cfg = Config{
		Chains:       []RawChainConfig{missingType},
		keystorePath: "",
	}

	err = cfg.validate()
	if err == nil {
		t.Fatal("must require type field")
	}

	cfg = Config{
		Chains:       []RawChainConfig{missingEndpoint},
		keystorePath: "",
	}

	err = cfg.validate()
	if err == nil {
		t.Fatal("must require endpoint field")
	}

	cfg = Config{
		Chains:       []RawChainConfig{missingName},
		keystorePath: "",
	}

	err = cfg.validate()
	if err == nil {
		t.Fatal("must require name field")
	}
}
