package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var (
	hdcontainsMnemonic string
	hdcontainsAccount  string
)

var hdcontainsCmd = &cobra.Command{
	Use:   "contains",
	Short: "if account is contained in wallet",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(hdcontainsMnemonic)
		if err != nil {
			log.Fatal(err)
		}
		address := common.HexToAddress(hdcontainsAccount)
		account := accounts.Account{
			Address: address,
		}
		is := wallet.Contains(account)
		fmt.Println(is)

	},
}

func init() {
	hdcontainsCmd.Flags().StringVarP(&hdcontainsMnemonic, "mnemonic", "m", "", "input mnemonic")
	hdcontainsCmd.MarkFlagRequired("mnemonic")

	hdcontainsCmd.Flags().StringVarP(&hdcontainsAccount, "account", "a", "", "input account")
	hdcontainsCmd.MarkFlagRequired("account")

	hdwalletCmd.AddCommand(hdcontainsCmd)
}
