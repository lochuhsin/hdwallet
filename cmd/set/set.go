/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package set

import (
	"github.com/spf13/cobra"
)

// SetCmd represents the set command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	SetCmd.AddCommand(networkCmd)
	SetCmd.AddCommand(mneCmd)
	SetCmd.AddCommand(passwordCmd)
	SetCmd.AddCommand(privatekeyCmd)
}
