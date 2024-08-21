/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "use for creating something, i.e wallet new [-m]",
	Long:  `...`,
	PreRun: func(cmd *cobra.Command, args []string) {
		pkg.InitWalletManager()
	},
	Run: func(cmd *cobra.Command, args []string) {
		config := service.GetConfig()

		coinSym, _ := pkg.CoinSelector("eth")
		wallet, err := pkg.GetWalletManager().NewWallet(coinSym, pkg.SetMnemonic(config.Mnemonic), pkg.SetPassword(config.Password))
		pk, _ := wallet.NewPrivateKey()
		_, err = crypto.HexToECDSA(pk)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {

}
