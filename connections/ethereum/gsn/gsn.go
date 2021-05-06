package gsn

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	Error  *RPCError       `json:"error"`
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
}

type gsnResponse struct {
	Fast          int64       `json:"fast"`
	Fastest       int64       `json:"fastest"`
	SafeLow       int64       `json:"safeLow"`
	Average       int64       `json:"average"`
	Block_time    float32     `json:"block_time"`
	BlockNum      int64       `json:"blockNum"`
	Speed         float32     `json:"speed"`
	SafeLowWait   float32     `json:"safeLowWait"`
	AvgWait       float32     `json:"avgWait"`
	FastWait      float32     `json:"fastWait"`
	FastestWait   float32     `json:"fastestWait"`
	GasPriceRange interface{} `json:"gasPriceRange"`
}

// callGSN will call the Gas Station Network and request the gas prices on the Ethereum network
func CallGSN(apiKey, speed string) (error, *big.Int) {
	time.Sleep(1 * time.Second)

	var gsnURL = "https://ethgasstation.info/api/ethgasAPI.json?api-key=" + apiKey

	var gsnRes gsnResponse

	retries := 10
	for i := 0; i < retries; i++ {
		res, err := http.Get(gsnURL)
		if err != nil {
			return err, nil
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err, nil
		}
		err = json.Unmarshal(body, &gsnRes)
		if err != nil {
			return err, nil
		}

		if body != nil && res.StatusCode == http.StatusOK {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil, getPrice(&gsnRes, speed)
}

func getPrice(result *gsnResponse, speed string) *big.Int {
	switch speed {
	case "fast":
		return big.NewInt(result.Fast)
	case "fastest":
		return big.NewInt(result.Fastest)
	case "average":
		return big.NewInt(result.Average)
	default:
		return big.NewInt(result.Fast)
	}
}
