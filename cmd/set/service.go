package set

import (
	"fmt"
	"wallet/pkg"
)

func getConfig() pkg.Config {
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
	return config
}
