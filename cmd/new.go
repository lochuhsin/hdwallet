/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var MISSING_FLAG_ERR = errors.New("missing required flags or empty value")

func createMnemonic() string {
	mn, err := pkg.NewMnemonic()
	if err != nil {
		fmt.Printf("Unable to generate mnemonic: %s", err)

	}
	fmt.Println("mnemonic:")
	fmt.Printf("%s \n", mn)
	return mn
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "use for creating something, i.e wallet new [-m]",
	Long:  `...`,
	PreRun: func(cmd *cobra.Command, args []string) {
		pkg.InitWalletManager()
	},
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		createMN, err := flags.GetBool("mnemonic")
		if err != nil {
			fmt.Println(err)
		}

		coin, err := flags.GetString("coin")
		if err != nil {
			fmt.Println(err)
		}

		if createMN {
			createMnemonic()
		}

		if coin != "" {
			config := service.GetConfig()
			if config.Mnemonic == "" {
				fmt.Println("no mnemonic found, create new one")
				mn := createMnemonic()
				config.Mnemonic = mn
			}

			coinSym, err := pkg.CoinSelector(coin)
			if err != nil {
				fmt.Println(err)
				return
			}
			wallet, err := pkg.GetWalletManager().NewWallet(coinSym, pkg.SetMnemonic(config.Mnemonic), pkg.SetPassword(config.Password))
			if err != nil {
				fmt.Println(err)
				return
			}
			pk, err := wallet.NewPrivateKey()
			symConfig := config.Symbols[coinSym]
			symConfig.PrivateKeys = []string{pk}
			config.Symbols[coinSym] = symConfig
			service.WriteConfig(config)
		}
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("mnemonic", "m", false, "Create a new mnemonic")
	newCmd.Flags().BoolP("supportWord", "s", false, "Create support word for mnemonic")
	newCmd.Flags().StringP("coin", "c", "", "Create a new wallet with mnemonic in config, if empty mnemonic is given, generate a new one")
}
