/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package show

import (
	"github.com/spf13/cobra"
)

// ShowCmd represents the set command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ShowCmd.AddCommand(showConfigCmd)
}
