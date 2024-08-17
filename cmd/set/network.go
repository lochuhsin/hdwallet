package set

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
import (
	"errors"
	"fmt"
	"strings"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var FlagMissingError = errors.New("missing required flags or empty value")

func isValidCoinHost(coin, host string) error {
	if coin == "" || host == "" {
		return FlagMissingError
	}
	upper := strings.ToUpper(coin)
	symbol, err := pkg.CoinSelector(upper)
	if err != nil {
		return err
	}
	_, err = pkg.GetClientStorage().GetClient(symbol, host)
	if err != nil {
		return err
	}
	return nil
}

// setCmd represents the set command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "A brief description of your command",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		pkg.InitClientStorage()
	},
	Run: func(cmd *cobra.Command, args []string) {
		coin, err := cmd.Flags().GetString("coin")
		if err != nil {
			fmt.Println(err)
			return
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := isValidCoinHost(coin, host); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	networkCmd.Flags().StringP("coin", "c", "", "The input coin symbol for setting network")
	networkCmd.Flags().StringP("host", "s", "", "The host of corresponding network")
}
