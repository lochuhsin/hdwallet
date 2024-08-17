package pkg

import (
	"fmt"
	"sync"
)

var clientSt *clientManager
var sOnce sync.Once // guard the initialization of storage

type IClient interface{}

type clientManager struct {
	storage map[CoinSymbol]IClient
}

func (c *clientManager) GetClient(type_ CoinSymbol, network string) (IClient, error) {
	cli, ok := c.storage[type_]
	if !ok {
		client, err := c.initClient(type_, network)
		c.storage[type_] = client
		if err != nil {
			return nil, err
		}
		cli = client
	}
	return cli, nil
}

func (c clientManager) initClient(type_ CoinSymbol, network string) (IClient, error) {
	switch type_ {
	case ETH:
		return initEthClient(network)
	default:
		return nil, fmt.Errorf("Unsupported CoinType: %v", type_)
	}
}

func newClientStorage() *clientManager {
	return &clientManager{
		storage: make(map[CoinSymbol]IClient),
	}
}

func InitClientStorage() {
	sOnce.Do(func() {
		clientSt = newClientStorage()
	})
}

func GetClientStorage() *clientManager {
	return clientSt
}
