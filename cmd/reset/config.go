package reset

import (
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "",
	Long:    ``,
	Aliases: []string{"c", "conf"},
	Run: func(cmd *cobra.Command, args []string) {
		defer fmt.Println("config reset complete")

		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
		}

		if coin != "" {
			config := service.GetConfig()
			sym, err := pkg.CoinSelector(coin)
			if err != nil {
				fmt.Println(err)
				return
			}
			config.Symbols[sym] = service.NewSymbolConfig()
			service.WriteConfig(config)
			return
		}
		conf := service.NewConfig()
		service.WriteConfig(conf)
	},
}

func init() {
	configCmd.Flags().StringP("coin", "c", "", "wallet coin type")
}
