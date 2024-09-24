# A simple Go chat app

A simple, TCP-based chat application implemented in Go. It features both client and server modes, allowing multiple clients to connect, send messages, and interact with each other in real-time. The application supports private messaging between users and basic commands like listing users or exiting the session.

## Features

- **Client-Server Architecture**: Clients connect to a central server and can send messages to each other.
- **Private Messaging**: Send messages directly to other users using `@username` format.
- **Basic Commands**: List available users, get help, or exit the chat.
- **Thread-Safe Client Management**: The server safely handles multiple concurrent clients.

## Getting Started

### Prerequisites

- **Go**: Ensure you have Go installed. This project uses Go version 1.22.7. You can download it from the official [Go website](https://golang.org/dl/).

#### Fedora

```bash
sudo dnf install go
```

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fuseraft/gochat.git
   cd gochat
   ```

2. Build the project:
   ```bash
   go build -o gochat
   ```

### Usage

You can run the application in two modes: `server` or `client`.

#### Starting the Server

1. Start the server by specifying the address and port:
   ```bash
   ./gochat --server <host-address:port>
   ```

   Example:
   ```bash
   ./gochat --server 127.0.0.1:8080
   ```

2. The server will listen for incoming connections from clients.

#### Starting the Client

1. Connect a client to the server by specifying the address and port of the running server:
   ```bash
   ./gochat --client <host-address:port>
   ```

   Example:
   ```bash
   ./gochat --client 127.0.0.1:8080
   ```

2. Once connected, you will be prompted to enter a username:
   ```bash
   username> yourname
   ```

3. You can now start sending messages or using commands.

### Commands

- **@username message**: Send a private message to `username`.
- **@all message**: Send a message to every connected user.
- **users**: List all connected users.
- **help**: Display a help message with available commands.
- **exit**: Exit the chat session.

### Example

1. **Running the server**:
   ```bash
   ./gochat --server localhost:12345
   ```
   Output:
   ```
   Server is listening on localhost:12345...
   ```

2. **Connecting a client**:
   ```bash
   ./gochat --client localhost:12345
   ```
   Output:
   ```
   Connected to server.
   username> alice
   > 
   ```

3. **Sending a message**:
   ```bash
   @bob Hello Bob!
   ```

4. **Listing users**:
   ```bash
   users
   ```

   Output:
   ```
   users: alice, bob
   ```

5. **Exiting the chat**:
   ```bash
   exit
   ```
