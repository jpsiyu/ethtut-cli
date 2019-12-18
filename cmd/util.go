package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jpsiyu/ethtut-cli/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	utAmount string
	utPriv   string
	utData   string
)

var utilCmd = &cobra.Command{
	Use:   "util",
	Short: "util operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var toWeiCmd = &cobra.Command{
	Use:   "towei",
	Short: "convert eth to wei",
	Run: func(cmd *cobra.Command, args []string) {
		wei := util.ToWei(utAmount, 18)
		fmt.Println("wei:", wei)
	},
}

var toEthCmd = &cobra.Command{
	Use:   "toeth",
	Short: "convert wei to eth",
	Run: func(cmd *cobra.Command, args []string) {
		eth := util.ToDecimal(utAmount, 18)
		fmt.Println("eth:", eth)
	},
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "generate signature",
	Run: func(cmd *cobra.Command, args []string) {
		privKey, err := crypto.HexToECDSA(utPriv)
		if err != nil {
			log.Fatal(err)
		}
		data := []byte(utData)
		hash := crypto.Keccak256Hash(data)
		signature, err := crypto.Sign(hash.Bytes(), privKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("hash:", hash.Hex())
		fmt.Println("signature:", hexutil.Encode(signature))
	},
}

func init() {
	rootCmd.AddCommand(utilCmd)

	utilCmd.AddCommand(toWeiCmd)
	toWeiCmd.Flags().StringVarP(&utAmount, "amount", "a", "", "amount")
	toWeiCmd.MarkFlagRequired("amount")

	utilCmd.AddCommand(toEthCmd)
	toEthCmd.Flags().StringVarP(&utAmount, "amount", "a", "", "amount")
	toEthCmd.MarkFlagRequired("amount")

	utilCmd.AddCommand(signCmd)
	signCmd.Flags().StringVarP(&utPriv, "priv", "p", "", "private key")
	signCmd.Flags().StringVarP(&utData, "data", "d", "", "data used to sign")
	utilCmd.MarkFlagRequired("priv")
	utilCmd.MarkFlagRequired("data")
}
