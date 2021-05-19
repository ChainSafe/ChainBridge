// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package egs

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleResponse = &gasPriceResponse{
	Fast:    1,
	Fastest: 2,
	Average: 4,
}

// Captured by querying endpoint manually
var rawResponse = "{\"fast\":590,\"fastest\":700,\"safeLow\":490,\"average\":0.5,\"block_time\":15.9,\"blockNum\":12389059,\"speed\":0.995538884040233,\"safeLowWait\":13.6,\"avgWait\":2.1,\"fastWait\":0.9,\"fastestWait\":0.6,\"gasPriceRange\":{\"4\":265,\"6\":265,\"8\":265,\"10\":265,\"20\":265,\"30\":265,\"40\":265,\"50\":265,\"60\":265,\"70\":265,\"80\":265,\"90\":265,\"100\":265,\"110\":265,\"120\":265,\"130\":265,\"140\":265,\"150\":265,\"160\":265,\"170\":265,\"180\":265,\"190\":265,\"200\":265,\"220\":265,\"240\":265,\"260\":265,\"280\":265,\"300\":265,\"320\":265,\"340\":265,\"360\":265,\"380\":265,\"400\":265,\"420\":265,\"440\":265,\"460\":265,\"480\":15,\"490\":13.6,\"500\":2.1,\"520\":1.7,\"540\":1.2,\"560\":1,\"580\":0.9,\"590\":0.9,\"600\":0.8,\"620\":0.7,\"640\":0.7,\"660\":0.7,\"680\":0.7,\"700\":0.6}}"

func TestParsePrice(t *testing.T) {
	assert.Equal(t, parsePrice(exampleResponse, Fastest), big.NewInt(200000000))
	assert.Equal(t, parsePrice(exampleResponse, Fast), big.NewInt(100000000))
	assert.Equal(t, parsePrice(exampleResponse, Average), big.NewInt(400000000))
}

func TestFetchPrice(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte(rawResponse))
	}))
	// Close the server when test finishes
	defer server.Close()

	res, err := queryAPI(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	var expected gasPriceResponse
	err = json.Unmarshal([]byte(rawResponse), &expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, *res, expected)
	assert.Equal(t, parsePrice(res, Fastest), big.NewInt(70000000000))
}
