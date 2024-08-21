package reset

import "github.com/spf13/cobra"

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ResetCmd.AddCommand(configCmd)
}
