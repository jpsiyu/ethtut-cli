package cmd

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
)

var (
	hdBits     int
	hdMnemonic string
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

func init() {
	rootCmd.AddCommand(hdwalletCmd)

	hdwalletCmd.AddCommand(hdgenerateCmd)
	hdgenerateCmd.Flags().IntVarP(&hdBits, "hdBits", "b", 256, "set mnemonic's length")

	hdwalletCmd.AddCommand(hdstatusCmd)
	hdstatusCmd.Flags().StringVarP(&hdMnemonic, "mnemonic", "m", "", "input mnemonic")
	hdstatusCmd.MarkFlagRequired("mnemonic")

}
