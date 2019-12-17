package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "whisper chat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chat called")
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
