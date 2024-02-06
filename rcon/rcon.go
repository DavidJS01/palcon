package rcon

import (
	"fmt"
	"log"

	"github.com/gorcon/rcon"
)

func ConnectToServer(serverHost string, port string, adminPassword string) *rcon.Conn {
	connString := fmt.Sprintf("%s:%s", serverHost, port)
	conn, err := rcon.Dial(connString, adminPassword)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func ExecuteCommand(command string, connection *rcon.Conn) (string, error) {
	response, err := connection.Execute(command)
	if err != nil {
		return "", fmt.Errorf("error executing command: %w", err)
	}
	return response, nil
}
