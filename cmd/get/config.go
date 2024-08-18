package get

import (
	"fmt"
	"wallet/cmd/service"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ShowCmd represents the set command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := service.ReadConfig()
		if err != nil {
			fmt.Printf("Unable to read config: %s \n", err)
			fmt.Println("Setting up default config")
			config = service.NewConfig()
			err := service.WriteConfig(config)
			if err != nil {
				fmt.Println(err)
			}
		}
		b, err := yaml.Marshal(config)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(string(b))
	},
}

func init() {
}
