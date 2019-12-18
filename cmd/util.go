package cmd

import (
	"fmt"
	"github.com/jpsiyu/ethtut-cli/util"
	"github.com/spf13/cobra"
)

var (
	utAmount string
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

func init() {
	rootCmd.AddCommand(utilCmd)

	utilCmd.AddCommand(toWeiCmd)
	toWeiCmd.Flags().StringVarP(&utAmount, "amount", "a", "", "amount")
	toWeiCmd.MarkFlagRequired("amount")

	utilCmd.AddCommand(toEthCmd)
	toEthCmd.Flags().StringVarP(&utAmount, "amount", "a", "", "amount")
	toEthCmd.MarkFlagRequired("amount")
}
