package get

import (
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

func validateCoinConfig(config service.SymbolConfig) bool {
	flag := true
	if config.Network == "" {
		flag = false
		fmt.Println("missing network settings")
	}
	if config.PrivateKeys == nil || len(config.PrivateKeys) == 0 {
		flag = false
		fmt.Println("missing or empty private keys settings")
	}
	if config.PrivateKeys == nil || len(config.PrivateKeys) == 0 {
		flag = false
		fmt.Println("missing or empty private keys settings")
	}
	return flag
}

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		pkg.InitClientStorage()
		pkg.InitWalletManager()
	},
	Run: func(cmd *cobra.Command, args []string) {
		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
			return
		}
		coinSym, err := pkg.CoinSelector(coin)
		if err != nil {
			fmt.Println(coinSym)
			return
		}
		config, err := service.ReadConfig()
		if err != nil {
			fmt.Println(coinSym)
			return
		}
		symConfig, ok := config.Symbols[coinSym]
		if !ok {
			fmt.Printf("missing %s coin configuration", coinSym)
		}
		if !validateCoinConfig(symConfig) {
			return
		}
		_, err = pkg.GetWalletManager().NewWallet(coinSym, pkg.SetMnemonic(config.Mnemonic), pkg.SetPassword(config.Password), pkg.SetPrivateKeys(symConfig.PrivateKeys), pkg.SetNetwork(symConfig.Network))
		if err != nil {
			fmt.Println(err)
			return
		}
		balance, err := pkg.GetBalance(coinSym)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("your balance %v \n", balance.String())
	},
}

func init() {
	balanceCmd.Flags().StringP("coin", "c", "", "Get balance from selected coin symbol")
}
