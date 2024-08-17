package pkg

import "github.com/ethereum/go-ethereum/ethclient"

type etherClient struct {
	client *ethclient.Client
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
