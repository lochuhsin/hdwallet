package pkg

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type etherClient struct {
	client *ethclient.Client
}

func (e *etherClient) GetBalance(address []byte) (*big.Int, error) {
	addr := common.BytesToAddress(address)
	return e.client.BalanceAt(context.Background(), addr, nil)
}

func initEthClient(network string) (*etherClient, error) {
	cli, err := ethclient.Dial(network)
	if err != nil {
		return nil, err
	}
	return &etherClient{
		client: cli,
	}, nil
}
