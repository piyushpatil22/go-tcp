package chatroom

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Client struct {
	name       string
	connection net.Conn
	hasNameSet bool
	server_id  string
}

func (Client *Client) HandleConn(server_list *[]Server) {
	reader := bufio.NewReader(Client.connection)
	if !Client.hasNameSet && strings.HasPrefix(Client.name, "guest") {
		Client.connection.Write([]byte("What is your name?"))
		new_name, err := reader.ReadString('\n')
		if err != nil {
			Client.connection.Close()
		}
		new_name = strings.Trim(new_name, "\n")
		fmt.Println(Client.name + " has changed name to " + new_name)
		Client.name = new_name
		Client.hasNameSet = true
	}
	for Client.server_id == "" {
		Client.connection.Write(
			[]byte(
				"Below is a list of available chat rooms you can join.\nEnter the number of the chat room you would like to join",
			),
		)
		for _, server := range *server_list {
			Client.connection.Write(
				[]byte(
					"\n" + strconv.FormatInt(
						int64(server.server_id),
						10,
					) + ": " + server.name + "\n",
				),
			)
		}
		server_id, err := reader.ReadString('\n')
		if err != nil {
			Client.connection.Write([]byte("Did not receive a valid server number"))
		}
		// check if input is valid and assign that id to the client
		for _, server := range *server_list {
			fmt.Println(server)
			parser_server_id, err := strconv.ParseInt(trimInput(&server_id), 10, 64)
			if err != nil {
				Client.connection.Close()
			}
			if server.server_id == int(parser_server_id) {
				Client.server_id = server_id
			}
		}
	}
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(Client.name + " disconnected")
			Client.connection.Close()
			return
		}
		fmt.Println(Client.name + ": " + message)
	}
}

func NewClient(name string, conn *net.Conn) *Client {
	return &Client{
		name:       name,
		connection: *conn,
	}
}
