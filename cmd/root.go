package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pi-connect",
	Short: "Posts and gets the ip address of the hosts to mongodb.",
	Long:  `Posts and gets the ip address of the hosts to mongodb. Can be used to find the local ip address of raspberry pi in headless setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
