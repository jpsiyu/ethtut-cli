package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jpsiyu/ethtut-cli/conf"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client info",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial(conf.ShhUrl)
		if err != nil {
			log.Fatal(err)
		}
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("client is visiting blockchain which chain id is", chainID)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
