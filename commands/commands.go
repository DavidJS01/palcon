package commands

import (
	"fmt"
	"strings"
)

type Command string

const (
	Save             Command = "Save"
	Exit             Command = "Shutdown"
	Start            Command = "Startup"
	Broadcast        Command = "Broadcast"
	KickPlayer       Command = "KickPlayer"
	BanPlayer        Command = "BanPlayer"
	TeleportToPlayer Command = "TeleportToPlayer"
	TeleportToMe     Command = "TeleportToMe"
	ShowPlayers      Command = "ShowPlayers"
	Info             Command = "Info"
)

func BuildCommand(command Command, args []string, seperator string) string {
	commandArgs := strings.Join(args, seperator)
	return fmt.Sprintf("%s %s", command, commandArgs)
}
