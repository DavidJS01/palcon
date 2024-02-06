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

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the palworld server",
	Run: func(cmd *cobra.Command, args []string) {
		host, port, password := GetConnectionInfo(rootCmd)
		conn := rcon.ConnectToServer(host, port, password)
		defer conn.Close()

		output, err := rcon.ExecuteCommand(string(commands.Save), conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Successfully saved! %s", output))
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
