package cmd

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var bits int

var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		mnemonic, err := hdwallet.NewMnemonic(bits)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mnemonic)
	},
}

func init() {
	mnemonicCmd.Flags().IntVarP(&bits, "bits", "b", 256, "set mnemonic's length")

	hdwalletCmd.AddCommand(mnemonicCmd)
}
