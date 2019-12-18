package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/spf13/cobra"
)

var (
	trBlock  int64
	trTxHash string
)

func displayBlock(block *types.Block) {
	tm := time.Unix(int64(block.Time()), 0)
	fmt.Println("number:", block.Number().Uint64())
	fmt.Println("time:", tm)
	fmt.Println("diffculty:", block.Difficulty().Uint64())
	fmt.Println("hash:", block.Hash().Hex())
	fmt.Println("transaction length:", len(block.Transactions()))
}

func displayTx(tx *types.Transaction) {
	fmt.Println("hash:", tx.Hash().Hex())
	fmt.Println("value:", tx.Value().String())
	fmt.Println("gas", tx.Gas())
	fmt.Println("nonce", tx.Nonce())
	fmt.Println("data", tx.Data())
	fmt.Println("to", tx.To().Hex())
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

var trTxCmd = &cobra.Command{
	Use:   "tx",
	Short: "query transaction",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}
		hash := common.HexToHash(trTxHash)
		tx, isPending, err := client.TransactionByHash(context.Background(), hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pending:", isPending)
		displayTx(tx)
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.AddCommand(trBlockCmd)
	trBlockCmd.Flags().Int64VarP(&trBlock, "block", "b", 0, "block number")

	transactionCmd.AddCommand(trTxCmd)
	trTxCmd.Flags().StringVarP(&trTxHash, "hash", "x", "", "input tx hash")
	trTxCmd.MarkFlagRequired("hash")
}
