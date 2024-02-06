/*
Copyright © 2024 David Shipman chemdev.me@protonmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "palcon",
	Long: `Palcon is a CLI tool that lets you interact with dedicated palworld servers.
	Example usage: palcon --host 173.194.0.0 --port 25575 --password PASSWORD info`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	serverHost := "127.0.0.1"
	rootCmd.PersistentFlags().StringVar(&serverHost, "host", "127.0.0.1", "IP of the dedicated server")

	serverPort := "25575"
	rootCmd.PersistentFlags().StringVar(&serverPort, "port", "25575", "Port of the dedicated server")

	adminPassword := "password"
	rootCmd.PersistentFlags().StringVar(&adminPassword, "password", "password", "Admin password of the dedicated server")

	serverExecutablePath := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PalServer\\PalServer.exe"
	rootCmd.PersistentFlags().StringVar(&serverExecutablePath, "serverExecutablePath", serverExecutablePath, "Path to Palworld dedicated server")

}

func GetConnectionInfo(*cobra.Command) (string, string, string) {
	host := rootCmd.PersistentFlags().Lookup("host").Value.String()
	port := rootCmd.PersistentFlags().Lookup("port").Value.String()
	password := rootCmd.PersistentFlags().Lookup("password").Value.String()

	return host, port, password

}
