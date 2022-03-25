package cmd

import (
	"github.com/andylolu2/pi-connect/pkg/hostinfo"
	"github.com/andylolu2/pi-connect/pkg/mongodb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(postIp)
}

var postIp = &cobra.Command{
	Use:   "postip",
	Short: "Post the host machine's ip address to a mongodb document.",
	Long:  `Post the host machine's local ip address to mongodb document in pi_connect/ip_addr together with the host's name.`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := hostinfo.GetHostName()
		hostIp := hostinfo.GetOutboundIP()

		mongodb.PostIp(hostname, hostIp)
	},
}
