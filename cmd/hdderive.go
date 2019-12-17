package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var (
	hdmnemonic string
	hdpath     string
)

var hdderiveCmd = &cobra.Command{
	Use:   "derive",
	Short: "derive account",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(hdmnemonic)
		if err != nil {
			log.Fatal(err)
		}
		derivePath, err := accounts.ParseDerivationPath(hdpath)
		if err != nil {
			log.Fatal(err)
		}
		account, err := wallet.Derive(derivePath, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(account.Address.Hex())
	},
}

func init() {
	hdderiveCmd.Flags().StringVarP(&hdmnemonic, "mnemonic", "m", "", "input mnemonic")
	hdderiveCmd.MarkFlagRequired("mnemonic")

	hdderiveCmd.Flags().StringVarP(&hdpath, "path", "p", "", "input derive path")
	hdderiveCmd.MarkFlagRequired("path")

	hdwalletCmd.AddCommand(hdderiveCmd)
}
