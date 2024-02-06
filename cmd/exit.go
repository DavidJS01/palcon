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

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Save the server and shut it down",
	Long: `
		Saves the server and then shuts down the server
		after 30 seconds.`,

	Run: func(cmd *cobra.Command, args []string) {
		host, port, password := GetConnectionInfo(rootCmd)
		conn := rcon.ConnectToServer(host, port, password)
		defer conn.Close()

		output, err := rcon.ExecuteCommand(string(commands.Save), conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Successfully saved! %s", output))

		exitCmd := commands.BuildCommand(commands.Exit, []string{"30"}, " ")
		output, err = rcon.ExecuteCommand(exitCmd, conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(exitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
