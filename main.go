package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		print_usage()
	}

	mode := os.Args[1]
	modes := map[string]bool{Mode.Server: true, Mode.Client: true}
	if !modes[mode] {
		print_usage()
	}

	address := os.Args[2]

	if mode == Mode.Server {
		run_server(address)
	} else {
		run_client(address)
	}
}

func print_usage() {
	fmt.Println("Usage: ./gochat --[server|client] <host-address:port>")
	os.Exit(0)
}
