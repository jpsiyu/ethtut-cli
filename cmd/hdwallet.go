package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hdwalletCmd = &cobra.Command{
	Use:   "hdwallet",
	Short: "hdwallet operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hdwallet called")
	},
}

func init() {
	rootCmd.AddCommand(hdwalletCmd)
}
