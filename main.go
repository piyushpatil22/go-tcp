package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/piyushpatil22/go-tcp/chatroom"
)

func main() {
	// create a TCP server on 8080 port
	guest_count := 1
	server, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// static server list

	var servers []chatroom.Server
	servers = append(servers, chatroom.NewChatRoom("valorant"))
	servers = append(servers, chatroom.NewChatRoom("counter strike"))
	fmt.Print("Server started at port 8080")
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		name := "guest" + strconv.FormatInt(int64(guest_count), 10)
		fmt.Println(name + " user joined!")
		client := chatroom.NewClient(name, &conn)

		go client.HandleConn(&servers)
		guest_count++
	}
}
