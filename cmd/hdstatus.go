package cmd

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var mnemonic string

var hdstatusCmd = &cobra.Command{
	Use:   "status",
	Short: "print wallet status",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(mnemonic)
		if err != nil {
			log.Fatal(err)
		}
		status, err := wallet.Status()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(status)
	},
}

func init() {
	hdstatusCmd.Flags().StringVarP(&mnemonic, "mnemonic", "m", "", "input mnemonic")
	hdstatusCmd.MarkFlagRequired("mnemonic")

	hdwalletCmd.AddCommand(hdstatusCmd)
}
