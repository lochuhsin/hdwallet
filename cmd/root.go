/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"wallet/cmd/create"
	"wallet/cmd/get"
	"wallet/cmd/reset"
	"wallet/cmd/set"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "wallet",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wallet.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.AddCommand(set.SetCmd)
	RootCmd.AddCommand(get.GetCmd)
	RootCmd.AddCommand(create.CreateCmd)
	RootCmd.AddCommand(testCmd)
	RootCmd.AddCommand(reset.ResetCmd)
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
