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
	hdBits     int
	hdMnemonic string
	hdPath     string
	hdAccount  string
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
		mnemonic, err := hdwallet.NewMnemonic(hdBits)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mnemonic)
	},
}

var hdstatusCmd = &cobra.Command{
	Use:   "status",
	Short: "print wallet status",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(hdMnemonic)
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

var hdderiveCmd = &cobra.Command{
	Use:   "derive",
	Short: "derive account",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(hdMnemonic)
		if err != nil {
			log.Fatal(err)
		}
		derivePath, err := accounts.ParseDerivationPath(hdPath)
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

var hdcontainsCmd = &cobra.Command{
	Use:   "contains",
	Short: "if account is contained in wallet",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := hdwallet.NewFromMnemonic(hdMnemonic)
		if err != nil {
			log.Fatal(err)
		}
		address := common.HexToAddress(hdAccount)
		account := accounts.Account{
			Address: address,
		}
		is := wallet.Contains(account)
		fmt.Println(is)

	},
}

func init() {
	rootCmd.AddCommand(hdwalletCmd)

	hdwalletCmd.AddCommand(hdgenerateCmd)
	hdgenerateCmd.Flags().IntVarP(&hdBits, "hdBits", "b", 256, "set mnemonic's length")

	hdwalletCmd.AddCommand(hdstatusCmd)
	hdstatusCmd.Flags().StringVarP(&hdMnemonic, "mnemonic", "m", "", "input mnemonic")
	hdstatusCmd.MarkFlagRequired("mnemonic")

	hdwalletCmd.AddCommand(hdderiveCmd)
	hdderiveCmd.Flags().StringVarP(&hdMnemonic, "mnemonic", "m", "", "input mnemonic")
	hdderiveCmd.MarkFlagRequired("mnemonic")
	hdderiveCmd.Flags().StringVarP(&hdPath, "path", "p", "", "input derive path")
	hdderiveCmd.MarkFlagRequired("path")

	hdwalletCmd.AddCommand(hdcontainsCmd)
	hdcontainsCmd.Flags().StringVarP(&hdMnemonic, "mnemonic", "m", "", "input mnemonic")
	hdcontainsCmd.MarkFlagRequired("mnemonic")
	hdcontainsCmd.Flags().StringVarP(&hdAccount, "account", "a", "", "input account")
	hdcontainsCmd.MarkFlagRequired("account")
}
