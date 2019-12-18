package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/spf13/cobra"
)

var (
	trBlock int64
)

func displayBlock(block *types.Block) {
	tm := time.Unix(int64(block.Time()), 0)
	fmt.Println("number:", block.Number().Uint64())
	fmt.Println("time:", tm)
	fmt.Println("diffculty:", block.Difficulty().Uint64())
	fmt.Println("hash:", block.Hash().Hex())
	fmt.Println("transaction length:", len(block.Transactions()))
}

var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "transaction operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var trBlockCmd = &cobra.Command{
	Use:   "block",
	Short: "query block info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}

		var number *big.Int = new(big.Int)
		if trBlock != 0 {
			number.SetInt64(trBlock)
		} else {
			number = nil
		}

		block, err := client.BlockByNumber(context.Background(), number)
		if err != nil {
			log.Fatal(err)
		}
		displayBlock(block)
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.AddCommand(trBlockCmd)
	trBlockCmd.Flags().Int64VarP(&trBlock, "block", "b", 0, "block number")

}
