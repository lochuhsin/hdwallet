package pkg

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethClient struct {
	client *ethclient.Client
}

func (e *ethClient) GetBalance(address []byte) (*big.Int, error) {
	addr := common.BytesToAddress(address)
	return e.client.BalanceAt(context.Background(), addr, nil)
}

func initEthClient(network string) (*ethClient, error) {
	cli, err := ethclient.Dial(network)
	if err != nil {
		return nil, err
	}
	return &ethClient{
		client: cli,
	}, nil
}
