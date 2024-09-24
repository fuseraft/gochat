package main

import "net"

type Client struct {
	ID   string
	Conn net.Conn
}

type Message struct {
	ClientID  string `json:"client_id"`
	Content   string `json:"content"`
	Recipient string `json:"recipient"`
}
