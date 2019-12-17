package cmd

import (
	"github.com/jpsiyu/ethtut-cli/chat"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "whisper chat",
	Run: func(cmd *cobra.Command, args []string) {
		chat.Run()
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
