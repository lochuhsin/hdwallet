package show

import (
	"fmt"
	"wallet/pkg"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// ShowCmd represents the set command
var showConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
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
