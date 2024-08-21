package create

import (
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var walletCmd = &cobra.Command{
	Use:     "wallet",
	Short:   "",
	Long:    ``,
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
			return
		}

		config := service.GetConfig()
		sym, err := pkg.CoinSelector(coin)
		if err != nil {
			fmt.Println(err)
			return
		}
		config.Symbols[sym] = service.NewSymbolConfig()
		service.WriteConfig(config)
	},
}

func init() {
	walletCmd.Flags().StringP("coin", "c", "", "wallet coin type")
}
