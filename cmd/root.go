package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	barkApi    string = os.Getenv("barkApi")

	rootCmd = &cobra.Command{
		Use:   "bark",
		Short: "A cli to use bark",
		Long:  `Bark is a CLI to send notification to ios`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	makeBarkUrl()

	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")

}
