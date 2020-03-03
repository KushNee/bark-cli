package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	title   string
	message string
)

var plainCommand = &cobra.Command{
	Use:   "plain",
	Short: "just send text to device",
	Run: func(cmd *cobra.Command, args []string) {
		if title != "" && message != "" {
			barkNotify(message, title)
		} else if title == "" && message != "" {
			barkNotify(message, "default")
		} else {
			fmt.Println("请至少输入发送内容")
		}
	},
}

func init() {
	rootCmd.AddCommand(plainCommand)
	plainCommand.Flags().StringVarP(&title, "title", "t", "", "title for notification")
	plainCommand.Flags().StringVarP(&message, "message", "m", "", "message for notification")

}
