package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	// create a TCP server on 8080 port
	guest_count := 1
	server, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Server started at port 8080")
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		name := "guest" + strconv.FormatInt(int64(guest_count), 10)
		fmt.Println(name + " user joined!")
		client := Client{
			name,
			conn,
			false,
		}
		go handleConn(client)
		guest_count++
	}
}

func handleConn(client Client) {
	reader := bufio.NewReader(client.connection)
	if !client.hasNameSet && strings.HasPrefix(client.name, "guest") {
		client.connection.Write([]byte("What is your name?"))
		new_name, err := reader.ReadString('\n')
		if err != nil {
			client.connection.Close()
		}
		new_name = strings.Trim(new_name, "\n")
		fmt.Println(client.name + " has changed name to " + new_name)
		client.name = new_name
		client.hasNameSet = true
	}

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(client.name + " disconnected")
			client.connection.Close()
			return
		}
		fmt.Println(client.name + ": " + message)
	}
}

type Client struct {
	name       string
	connection net.Conn
	hasNameSet bool
}
