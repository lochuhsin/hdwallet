package create

import (
	"fmt"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mn, err := pkg.NewMnemonic()
		if err != nil {
			fmt.Printf("Unable to generate mnemonic: %s", err)

		}
		fmt.Println("mnemonic:")
		fmt.Printf("%s \n", mn)
	},
}

func init() {
}
