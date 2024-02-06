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

// banCmd represents the banplayer command
var banCmd = &cobra.Command{
	Use:   "ban",
	Short: "Ban a player by player UID OR steam ID",
	Run: func(cmd *cobra.Command, args []string) {
		host, port, password := GetConnectionInfo(rootCmd)
		conn := rcon.ConnectToServer(host, port, password)
		defer conn.Close()

		command := commands.BuildCommand(commands.BanPlayer, args, " ")
		output, err := rcon.ExecuteCommand(command, conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(output)
	},
}

func init() {
	rootCmd.AddCommand(banCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// banplayerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// banplayerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
