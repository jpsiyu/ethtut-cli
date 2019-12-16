package cmd

import (
	"github.com/spf13/cobra"
)

var hdwalletCmd = &cobra.Command{
	Use:   "hdwallet",
	Short: "hdwallet operation",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(hdwalletCmd)
}
