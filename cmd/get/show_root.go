/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"github.com/spf13/cobra"
)

// GetCmd represents the set command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	GetCmd.AddCommand(configCmd)
	GetCmd.AddCommand(balanceCmd)
}
