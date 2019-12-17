package cmd

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/spf13/cobra"
)

var (
	accAccount string
	accEth     bool
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account operation",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "check account balance",
	Run: func(cmd *cobra.Command, args []string) {
		address := common.HexToAddress(accAccount)
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}
		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			log.Fatal(err)
		}

		if accEth != true {
			fmt.Println("balance:", balance, "wei")
		} else {
			fbalance := new(big.Float).SetInt(balance)
			ethValue := fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Println("balance:", ethValue.String(), "eth")
		}

	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(balanceCmd)
	balanceCmd.Flags().StringVarP(&accAccount, "account", "a", "", "account")
	balanceCmd.MarkFlagRequired("account")
	balanceCmd.Flags().BoolVarP(&accEth, "eth", "e", false, "if convert to eth")
}
