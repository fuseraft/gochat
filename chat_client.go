package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func run_client(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")
	fmt.Print("username> ")
	clientID, err := readInput()
	if err != nil {
		fmt.Println("Could not get user input:", err)
	}

	err = send(conn, clientID, "", "")
	if err != nil {
		fmt.Println("Failed to register username with server:", err)
		os.Exit(0)
	}

	go listenForMessages(conn)

	var command string

	for {
		prompt()
		command, err = readInput()
		if err != nil {
			fmt.Println("Could not get user input:", err)
			continue
		}
		parseCommand(conn, clientID, command)
	}
}

func parseCommand(conn net.Conn, clientID string, command string) {
	if strings.HasPrefix(command, "@") {
		parts := strings.SplitN(command, " ", 2)

		if len(parts) < 2 {
			return
		}

		recipient := parts[0]
		if recipient != "@all" {
			recipient = strings.TrimPrefix(recipient, "@")
		}
		content := parts[1]

		err := send(conn, clientID, content, recipient)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		} else {
			fmt.Printf("%s: %s\n", clientID, content)
		}
	} else if isCommand(command) {
		err := send(conn, clientID, command, "")
		if err != nil {
			fmt.Println("Failed to send message:", err)
		}
		if command == Command.Exit {
			os.Exit(0)
		}
	} else {
		prompt()
	}
}

func prompt() {
	fmt.Print("> ")
}

func send(conn net.Conn, clientID string, content string, recipient string) error {
	msg := Message{
		ClientID:  clientID,
		Content:   content,
		Recipient: recipient,
	}

	err := json.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(0)
	}

	return err
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func listenForMessages(conn net.Conn) {
	decoder := json.NewDecoder(conn)
	for {
		var msg Message
		err := decoder.Decode(&msg)

		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Connection closed by server.")
				os.Exit(0)
			} else {
				log.Println("Error while receiving server data:", err)
			}
			return
		}

		fmt.Printf("\n%s: %s\n", msg.ClientID, msg.Content)
		prompt()
	}
}
