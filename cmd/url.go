package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

var (
	url string
)

var urlCommand = &cobra.Command{
	Use:   "url",
	Short: "send url to device",
	Run: func(cmd *cobra.Command, args []string) {
		urlRegex, _ := regexp.Compile("http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*(),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+")
		if url != "" && urlRegex.MatchString(url) {
			if title != "" {
				barkUrl(url, title)
			} else {
				barkUrl(url, url)
			}
		} else {
			fmt.Println("请输入正确的 URL(带有协议名)")
		}
	},
}

func init() {
	rootCmd.AddCommand(urlCommand)
	urlCommand.Flags().StringVarP(&title, "title", "t", "", "title for notification")
	urlCommand.Flags().StringVarP(&url, "url", "u", "", "url for notification")
}
