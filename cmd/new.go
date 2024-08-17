/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"wallet/pkg"

	"github.com/spf13/cobra"
)

var MISSING_FLAG_ERR = errors.New("missing required flags or empty value")

func createMnemonic() {
	mn, err := pkg.NewMnemonic()
	if err != nil {
		fmt.Printf("Unable to generate mnemonic: %s", err)

	}
	fmt.Println("mnemonic:")
	fmt.Printf("%s \n", mn.MN)
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "use for creating something, i.e wallet new [-m]",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		createMN, err := flags.GetBool("mnemonic")
		if err != nil {
			fmt.Println(err)
		}
		if createMN {
			createMnemonic()
		}
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.Flags().BoolP("mnemonic", "m", false, "Create a new mnemonic")
}
