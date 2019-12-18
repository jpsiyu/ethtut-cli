package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/spf13/cobra"
)

var (
	trBlock  int64
	trTxHash string
	trPriv   string
	trTo     string
	trValue  int64
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

var trSendCmd = &cobra.Command{
	Use:   "send",
	Short: "send eth",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}
		privKey, err := crypto.HexToECDSA(trPriv)
		if err != nil {
			log.Fatal(err)
		}
		pubKey := privKey.Public()
		pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		value := big.NewInt(trValue)
		gasLimit := uint64(21000)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		toAddress := common.HexToAddress(trTo)
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
		if err != nil {
			log.Fatal(err)
		}
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tx send:", signedTx.Hash().Hex())
	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)

	transactionCmd.AddCommand(trBlockCmd)
	trBlockCmd.Flags().Int64VarP(&trBlock, "block", "b", 0, "block number")

	transactionCmd.AddCommand(trTxCmd)
	trTxCmd.Flags().StringVarP(&trTxHash, "hash", "x", "", "input tx hash")
	trTxCmd.MarkFlagRequired("hash")

	transactionCmd.AddCommand(trSendCmd)
	trSendCmd.Flags().StringVarP(&trPriv, "priv", "p", "", "private key")
	trSendCmd.Flags().StringVarP(&trTo, "to", "t", "", "to address")
	trSendCmd.Flags().Int64VarP(&trValue, "value", "v", 0, "value in wei")
	trSendCmd.MarkFlagRequired("priv")
	trSendCmd.MarkFlagRequired("to")
	trSendCmd.MarkFlagRequired("value")
}
