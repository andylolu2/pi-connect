package cmd

import (
	"fmt"

	"github.com/andylolu2/pi-connect/pkg/mongodb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getIp)
}

var getIp = &cobra.Command{
	Use:   "getip [hostname]",
	Short: "Get the specified host machine's ip address from a mongodb document.",
	Long:  "Get the specified host machine's ip address from a mongodb document.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]

		ip := mongodb.GetIp(hostname)
		fmt.Println(ip)
	},
}
