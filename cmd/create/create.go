package create

import (
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	CreateCmd.AddCommand(transactionCmd)
	CreateCmd.AddCommand(mnemonicCmd)
	CreateCmd.AddCommand(walletCmd)
}
