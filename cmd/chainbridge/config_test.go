package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/ChainSafe/ChainBridgeV2/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli"
)

func createTempConfigFile() (*os.File, *Config) {
	testConfig := NewConfig()
	ethCfg := core.ChainConfig{
		Endpoint: "",
		Home:     "",
		Away:     "",
		From:     "",
	}
	testConfig.Ethereum = ethCfg
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Crit("Cannot create temporary file", "err", err)
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

	if !reflect.DeepEqual(res.Ethereum, cfg.Ethereum) {
		t.Errorf("did not match\ngot: %+v\nexpected: %+v", res.Ethereum, cfg.Ethereum)
	}
}
