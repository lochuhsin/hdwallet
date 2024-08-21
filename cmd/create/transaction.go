package create

import (
	"fmt"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var transactionCmd = &cobra.Command{
	Use:     "transaction",
	Short:   "",
	Long:    ``,
	Aliases: []string{"trans"},
	Run: func(cmd *cobra.Command, args []string) {

		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
			return
		}
		to, err := cmd.Flags().GetString("to")
		if err != nil {
			fmt.Println(err)
			return
		}
		amount, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			fmt.Println(err)
			return
		}

		if coin == "" || to == "" || amount <= 0.0 {
			fmt.Println("all parameters are required, amount should be larger than zero")
			return
		}
		c, err := pkg.CoinSelector(coin)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = pkg.MakeTransaction(c, to, amount)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	transactionCmd.Flags().StringP("coin", "c", "", "the type of coin to make transaction")
	transactionCmd.Flags().StringP("to", "t", "", "the target address in hex")
	transactionCmd.Flags().Float64P("amount", "a", 0, "the amount of money to send")
}
