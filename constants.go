package main

var Mode = struct {
	Client string
	Server string
}{
	Client: "--client",
	Server: "--server",
}

var Command = struct {
	ListUsers string
	Exit      string
	Help      string
}{
	ListUsers: "users",
	Exit:      "exit",
	Help:      "help",
}

func isCommand(input string) bool {
	switch input {
	case Command.ListUsers, Command.Exit, Command.Help:
		return true
	default:
		return false
	}
}
