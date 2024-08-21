package pkg

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethClient struct {
	client *ethclient.Client
}

func (e *ethClient) GetBalance(address []byte) (*big.Float, error) {
	addr := common.BytesToAddress(address)
	return e.getBalance(addr)
}

func (e *ethClient) getBalance(addr common.Address) (*big.Float, error) {
	balance, err := e.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}
	fbalance := big.NewFloat(0)
	fbalance.SetString(balance.String())
	return WeiToEth(fbalance), nil
}

func (e *ethClient) SendTransaction(privateKeys []*ecdsa.PrivateKey, to *common.Address, amount float64) error {
	/**
	 * Calculate the amount of wei to send in each address
	 */
	fzero := big.NewFloat(0)
	ftotal := big.NewFloat(0)
	ftotal.SetFloat64(amount)

	emptyAccounts := make([]*ecdsa.PrivateKey, 0, len(privateKeys))
	for _, key := range privateKeys {
		fbalance, err := e.getBalance(crypto.PubkeyToAddress(key.PublicKey))
		if err != nil {
			return err
		}

		if fbalance.Cmp(fzero) == 0 {
			continue
		}

		diff := ftotal.Sub(ftotal, fbalance)
		if diff.Cmp(fzero) <= 0 {
			e.sendTransaction(key, to, ftotal)
			break
		}
		e.sendTransaction(key, to, fbalance)
		ftotal = diff
		emptyAccounts = append(emptyAccounts, key)
	}
	/**
	 * TODO: remove empty accounts, since they are not needed anymore
	 * emptyAccounts ....
	 */

	return nil
}

func (e *ethClient) sendTransaction(sender *ecdsa.PrivateKey, to *common.Address, amount *big.Float) error {
	public := sender.PublicKey
	senderAddr := crypto.PubkeyToAddress(public)
	nonce, err := e.client.PendingNonceAt(context.Background(), senderAddr)
	if err != nil {
		log.Fatal(err)
	}
	// set gas limit
	gasLimit := uint64(21000) // gas unit upper bound

	// get suggested gas price
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// convert eth to units of wei
	weiAmount := EthToWei(amount)
	tx := types.NewTransaction(nonce, *to, weiAmount, gasLimit, gasPrice, nil)

	// sign the transaction
	signedTx, err := e.signTransaction(tx, sender)
	if err != nil {
		return err
	}

	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	return nil
}

func (e *ethClient) signTransaction(tx *types.Transaction, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
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

func EthToWei(eth *big.Float) *big.Int {
	/**
	 * Do it better, like decimal or something
	 */
	weiUnit := big.NewFloat(1000000000000000000)
	val := eth.Mul(eth, weiUnit)
	res, _ := val.Int(nil)
	return res
}

func WeiToEth(wei *big.Float) *big.Float {
	return new(big.Float).Quo(wei, big.NewFloat(math.Pow10(18)))
}
