package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var clients = make(map[string]*Client)
var mutex = &sync.Mutex{}

func run_server(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
	defer listener.Close()
	log.Printf("Server is listening on %s...\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Could not establish connection:", err)
			continue
		}

		decoder := json.NewDecoder(conn)
		var initialMsg Message
		err = decoder.Decode(&initialMsg)

		if err != nil {
			log.Println("Invalid initial message from client:", err)
			conn.Close()
			continue
		}

		client := &Client{ID: initialMsg.ClientID, Conn: conn}

		mutex.Lock()
		clients[client.ID] = client
		mutex.Unlock()

		log.Printf("Client %s connected\n", client.ID)

		go handle(client)
	}
}

func broadcast(message Message) {
	mutex.Lock()
	defer mutex.Unlock()

	if message.Recipient == "@all" {
		for _, client := range clients {
			encoder := json.NewEncoder(client.Conn)
			err := encoder.Encode(message)
			if err != nil {
				log.Printf("Error sending message to client: %s\nError: %+v\n", client.ID, err)
			}
		}
		return
	}

	if recipient, ok := clients[message.Recipient]; ok {
		encoder := json.NewEncoder(recipient.Conn)
		err := encoder.Encode(message)
		if err != nil {
			log.Printf("Error sending message to recipient: %s\nError: %+v\n", message.Recipient, err)
		}
	} else {
		log.Printf("Unknown recipient: %s\n", message.Recipient)
		if sender, ok := clients[message.ClientID]; ok {
			send(sender.Conn, "server", fmt.Sprintf("User %s not found.", message.Recipient), sender.ID)
		} else {
			log.Printf("Failed to send error to unknown recipient: %s", message.ClientID)
		}
	}
}

func handle(client *Client) {
	defer func() {
		mutex.Lock()
		delete(clients, client.ID)
		mutex.Unlock()
		client.Conn.Close()
	}()

	decoder := json.NewDecoder((client.Conn))
	for {
		var message Message
		err := decoder.Decode(&message)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Client %s disconnected.\n", client.ID)
			} else {
				log.Println("Error decoding JSON:", err)
			}
			return
		}

		log.Printf("Received message from %s: %+v\n", client.ID, message)

		if isCommand(message.Content) {
			encoder := json.NewEncoder(client.Conn)
			switch message.Content {
			case Command.Help:
				encoder.Encode(getHelpMessage())
			case Command.ListUsers:
				encoder.Encode(getListUsersMessage())
			case Command.Exit:
				return
			}
			continue
		}

		broadcast(message)
	}
}

func createMessage(content string) Message {
	return Message{
		ClientID:  "",
		Content:   content,
		Recipient: "",
	}
}

func getHelpMessage() Message {
	return createMessage(`commands:
	help       show this message
	users      list users
	exit       exit the session`)
}

func getListUsersMessage() Message {
	var users []string

	mutex.Lock()
	for key := range clients {
		users = append(users, key)
	}
	mutex.Unlock()
	return createMessage(fmt.Sprintf("users: %s", strings.Join(users, ", ")))
}
