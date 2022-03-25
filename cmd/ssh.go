package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/andylolu2/pi-connect/pkg/mongodb"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ssh)
}

var ssh = &cobra.Command{
	Use:   "ssh [hostname]",
	Short: "Start ssh session to the specified host machine's name.",
	Long:  "Start ssh session to the specified host machine's name by retrieving its ip address from a mongodb document.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]

		ip := mongodb.GetIp(hostname)
		sshName := fmt.Sprintf("pi@%s", ip)

		command := exec.Command("ssh", "-o", "identitiesonly=yes", sshName)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			log.Fatalln(err)
		}
	},
}
