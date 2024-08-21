package set

import (
	"fmt"
	"wallet/cmd/service"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var privatekeyCmd = &cobra.Command{
	Use:     "privatekey",
	Short:   "",
	Long:    ``,
	Aliases: []string{"p"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Specify at least one argument to be password")
			return
		}

		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
			return
		}
		overwrite, err := cmd.Flags().GetBool("overwrite")
		if err != nil {
			fmt.Println(err)
			return
		}
		append_, err := cmd.Flags().GetBool("append")
		if err != nil {
			fmt.Println(err)
			return
		}

		if coin == "" {
			fmt.Println("Specify at least one coin type")
			return
		}
		sym, err := pkg.CoinSelector(coin)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !overwrite && !append_ {
			fmt.Println("Specify at least one operation, either overwrite or append, -o or -a")
		}

		if overwrite && append_ {
			fmt.Println("Unable to perform both overwrite and append in the same time")
		}
		config := service.GetConfig()
		obj := config.Symbols[sym]
		if overwrite {
			obj.PrivateKeys = args
		}

		if append_ {
			obj.PrivateKeys = append(obj.PrivateKeys, args...)
		}
		config.Symbols[sym] = obj
		service.WriteConfig(config)
	},
}

func init() {
	privatekeyCmd.Flags().StringP("coin", "c", "", "Specify a type of wallet")
	privatekeyCmd.Flags().BoolP("overwrite", "o", false, "when enabled, overwrite the private keys in wallet")
	privatekeyCmd.Flags().BoolP("append", "a", false, "when enabled append new keys to wallet")
}
