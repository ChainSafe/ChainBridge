package ethereum

import (
	"context"
	"testing"	
)

var TestEthereumEndpoint = "https://rinkeby.infura.io/v3/b0a01296903f4812b5ec2cf26cbded48"

func TestConnect(t *testing.T) {
	ctx := context.Background()
	conn := NewConnection(ctx, TestEthereumEndpoint)
	err := conn.Connect()
	if err != nil {
		t.Fatal(err)
	}
	conn.Close()
} 

func TestSendTx(t *testing.T) {
	
}