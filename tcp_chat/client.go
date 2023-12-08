package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"unicode"
)

type client struct {
	conn     net.Conn
	commands chan command
	server   *server
}

func (c *client) getInput() {
	for {
		input, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
				delete(c.server.clients, c.conn.RemoteAddr())
				c.server.broadcast(c, "left the server")
				c.conn.Close()
				return
			} else {
				log.Println(err)
				return
			}
		}

		input = clearString(strings.TrimSpace(input))
		args := strings.Split(input, " ")
		cmd := clearString(strings.TrimSpace(args[0]))

		log.Printf("%s enter input '%s' & cmd '%s'", c.conn.RemoteAddr().String(), input, cmd)

		switch cmd {
		case "/chat":
			c.commands <- command{
				ID:     CMD_CHAT,
				args:   args,
				client: c,
			}
		case "/quit":
			c.commands <- command{
				ID:     CMD_QUIT,
				args:   args,
				client: c,
			}
		default:
			c.sendError(fmt.Errorf("invalid command '%s', %s", cmd, availableCommandsMsg))
		}
	}
}

func (c *client) sendMessage(message string) {
	c.conn.Write([]byte(message + "\n\r"))
}

func (c *client) sendError(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n\r"))
}

func clearString(str string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsSymbol(r) {
			return r
		}
		return -1
	}, str)
}
