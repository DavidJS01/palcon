/*
Copyright © 2024 David Shipman chemdev.me@protonmail.com
*/
package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a server on this local computer by starting the Palworld server",
	Long:  `Start a server using the --serverExecutablePath flag's value`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServer() {
	executableCommand := exec.Command(rootCmd.PersistentFlags().Lookup("serverExecutablePath").Value.String())
	if err := executableCommand.Start(); err != nil {
		log.Fatal(err)
	}
}
