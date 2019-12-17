package cmd

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var bits int

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
	hdgenerateCmd.Flags().IntVarP(&bits, "bits", "b", 256, "set mnemonic's length")

	hdwalletCmd.AddCommand(hdgenerateCmd)
}
