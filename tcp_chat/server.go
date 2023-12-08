package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	clients  map[net.Addr]*client
	commands chan command
}

func newServer() *server {
	return &server{
		clients:  make(map[net.Addr]*client),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		commands: s.commands,
		server:   s,
	}

	s.clients[c.conn.RemoteAddr()] = c
	s.broadcast(c, "joined the server")
	c.sendMessage(fmt.Sprintf("Welcome to server :%v", port))

	c.getInput()
}

func (s *server) executeCommand() {
	for cmd := range s.commands {
		switch cmd.ID {
		case CMD_CHAT:
			s.chat(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		default:
			cmd.client.sendError(fmt.Errorf("invalid command, %s", availableCommandsMsg))
		}
	}
}

func (s *server) broadcast(c *client, message string) {
	for clientAddr, client := range s.clients {
		if clientAddr != c.conn.RemoteAddr() {
			client.sendMessage(c.conn.RemoteAddr().String() + ": " + message)
		}
	}
}

func (s *server) chat(c *client, args []string) {
	message := strings.Join(args[1:], " ")
	c.sendMessage("you: "+message)
	s.broadcast(c, message)
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())

	if c.server != nil {
		c.server = nil
		delete(s.clients, c.conn.RemoteAddr())
		s.broadcast(c, "left the server")
	}

	c.conn.Close()
}
