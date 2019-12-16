package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
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
	accountCmd.AddCommand(generateCmd)
}
