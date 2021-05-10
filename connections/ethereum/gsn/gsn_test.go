package gsn

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ExampleResponse = &gasPriceResponse{
	Fast:          1,
	Fastest:       2,
	SafeLow:       3,
	Average:       4,
	BlockTime:     5,
	BlockNum:      6,
	Speed:         7,
	SafeLowWait:   8,
	AvgWait:       9,
	FastWait:      10,
	FastestWait:   11,
	GasPriceRange: nil,
}

func TestParsePrice(t *testing.T) {
	assert.Equal(t, parsePrice(ExampleResponse, Fastest), big.NewInt(ExampleResponse.Fastest))
	assert.Equal(t, parsePrice(ExampleResponse, Fast), big.NewInt(ExampleResponse.Fast))
	assert.Equal(t, parsePrice(ExampleResponse, Average), big.NewInt(ExampleResponse.Average))
}

func TestFetchPrice(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		bz, err := json.Marshal(ExampleResponse)
		if err != nil {
			panic(err)
		}
		_, _ = rw.Write(bz)
	}))
	// Close the server when test finishes
	defer server.Close()

	res, err := queryAPI(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res, ExampleResponse)
	assert.Equal(t, parsePrice(ExampleResponse, Fastest), big.NewInt(ExampleResponse.Fastest))
}
