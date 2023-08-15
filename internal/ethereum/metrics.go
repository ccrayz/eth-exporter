package ethereum

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type MetricsCollector struct {
	ethClient *rpc.Client
	// Add more metric related variables here
}

func NewMetricsCollector(rpcEndpoint string) *MetricsCollector {
	ethClient, err := rpc.Dial(rpcEndpoint) // Replace with your Ethereum node's URL
	if err != nil {
		log.Fatal(err)
	}
	return &MetricsCollector{
		ethClient: ethClient,
	}
}

func (mc *MetricsCollector) GetAccountBalance(address string) (*big.Float, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid Ethereum address")
	}

	var balanceHex string
	err := mc.ethClient.Call(&balanceHex, "eth_getBalance", address, "latest")
	if err != nil {
		return nil, err
	}

	balanceWei, success := new(big.Int).SetString(balanceHex[2:], 16)
	if !success {
		return nil, fmt.Errorf("failed to parse balance")
	}

	// Convert Wei to Ether
	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balanceWei), big.NewFloat(1e18))
	return balanceEth, nil
}
