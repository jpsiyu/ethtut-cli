package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate account",
	Run: func(cmd *cobra.Command, args []string) {
		privKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		privKeyBytes := crypto.FromECDSA(privKey)
		privKeyHex := hexutil.Encode(privKeyBytes)
		fmt.Println("privkey", privKeyHex)

		pubKey := privKey.Public()
		pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		address := crypto.PubkeyToAddress(*pubKeyECDSA).Hex()
		fmt.Println("address:", address)
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(balanceCmd)
	balanceCmd.Flags().StringVarP(&accAccount, "account", "a", "", "account")
	balanceCmd.MarkFlagRequired("account")
	balanceCmd.Flags().BoolVarP(&accEth, "eth", "e", false, "if convert to eth")

	accountCmd.AddCommand(generateCmd)
}
