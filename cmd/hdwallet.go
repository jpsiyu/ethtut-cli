package cmd

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var (
	bits int
)

var hdwalletCmd = &cobra.Command{
	Use:   "hdwallet",
	Short: "hdwallet operation",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var hdgenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate mnemonic",
	Run: func(cmd *cobra.Command, args []string) {
		mnemonic, err := hdwallet.NewMnemonic(bits)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mnemonic)
	},
}

func init() {
	rootCmd.AddCommand(hdwalletCmd)

	hdwalletCmd.AddCommand(hdgenerateCmd)
	hdgenerateCmd.Flags().IntVarP(&bits, "bits", "b", 256, "set mnemonic's length")
}
