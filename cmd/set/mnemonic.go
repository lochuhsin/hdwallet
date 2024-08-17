package set

import (
	"errors"
	"fmt"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var MISSING_FLAG_ERR = errors.New("missing required flags or empty value")

var mneCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "use for creating something, i.e wallet new [-m]",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Specify at least one args for mnemonic")
			return
		}

		mn := args[0]

		config, err := pkg.ReadConfig()
		if err != nil {
			fmt.Printf("Unable to read config: %s \n", err)
			fmt.Println("Setting up default config")
			config = pkg.NewConfig()
			err := pkg.WriteConfig(config)
			if err != nil {
				fmt.Println(err)
			}
		}

		config.Mnemonic = mn
		err = pkg.WriteConfig(config)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
}
