package set

import (
	"fmt"
	"wallet/cmd/service"

	"github.com/spf13/cobra"
)

// GetCmd represents the set command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Specify at least one argument to be password")
			return
		}
		password := args[0]
		config := service.GetConfig()
		config.Password = password
		service.WriteConfig(config)
	},
}

func init() {
}
