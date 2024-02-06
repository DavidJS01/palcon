/*
Copyright © 2024 David Shipman chemdev.me@protonmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/DavidJS01/palcon/commands"
	"github.com/DavidJS01/palcon/rcon"
	"github.com/spf13/cobra"
)

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "broadcast a message to the server",
	Long: `
	Broadcast a message to the server.
	Currently, Palworld does not support spaces between words in the message, so
	the message is sent in snake case. Example: "hello_world"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		host, port, password := GetConnectionInfo(rootCmd)
		conn := rcon.ConnectToServer(host, port, password)
		defer conn.Close()

		broadcastCommand := commands.BuildCommand(commands.Broadcast, args, "_")
		output, err := rcon.ExecuteCommand(broadcastCommand, conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(output)
	},
}

func init() {
	rootCmd.AddCommand(broadcastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
