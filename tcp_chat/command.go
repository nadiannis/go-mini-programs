package main

const (
	CMD_CHAT int = iota
	CMD_QUIT
)

const availableCommandsMsg = "available commands: /chat [message] | /quit"

type command struct {
	ID     int
	args   []string
	client *client
}
