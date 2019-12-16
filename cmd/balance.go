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

var account string
var eth bool

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "check account balance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("balance of", account)
		address := common.HexToAddress(account)
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}
		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			log.Fatal(err)
		}

		if eth != true {
			fmt.Println("balance:", balance, "wei")
		} else {
			fbalance := new(big.Float).SetInt(balance)
			ethValue := fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
			fmt.Println("balance:", ethValue, "eth")
		}

	},
}

func init() {
	balanceCmd.Flags().StringVarP(&account, "account", "a", "", "account")
	balanceCmd.MarkFlagRequired("account")

	balanceCmd.Flags().BoolVarP(&eth, "eth", "e", false, "if convert to eth")

	accountCmd.AddCommand(balanceCmd)
}
