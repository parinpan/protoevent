
# ProtoEvent
ProtoEvent is an event-based TCP/UDP connection handling in Golang. It's simple and not exploiting too much the basic functionality of making a new protocol connection. Basically, ProtoEvent reimplements `net.Listener` and `net.Conn`  interface to have extended ability of capturing various events happened in network communication.

# Background
It's painful enough when you have to look after your simple network application codebase which is not well-organized ends up with producing more spagetthi code.

# Installation
```
go get -u https://github.com/parinpan/protoevent
```

# Demo
Here will be a demo.

# Examples
Server side application:
```golang
package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/parinpan/protoevent"
)

type Message struct {
	From string `json:"from"`
}

func main() {
	servant, event, err := protoevent.CreateServant("tcp", "0.0.0.0:8089")

	if nil != err {
		panic(err)
	}

	event.OnConnectionError(func(err error) {
	
	})

	event.OnConnectionAccepted(func(conn net.Conn) {
		fmt.Println("Accepting new connection: ", conn.RemoteAddr())
	})

	event.OnConnectionClosed(func(conn net.Conn) {
	
	})

	event.OnReceiveMessageError(func(conn net.Conn, err error) {
	
	})

	event.OnMessageReceived(func(conn net.Conn, message []byte) {
		fmt.Println("Received a message: ", string(message))
		
		var msg Message
		_ = json.Unmarshal(message, &msg)
		
		// send a message back to client
		sayHiMessage := fmt.Sprint("Hi ", msg.From, ". Welcome to ProtoEvent!")
		conn.Write([]byte(sayHiMessage))
	})

	event.OnSendMessageError(func(conn net.Conn, message []byte, err error) {
	
	})

	event.OnMessageSent(func(conn net.Conn, message []byte) {
		fmt.Println("Sent a message: ", string(message))
	})

	servant.Serve()
}
```
---
Client side application:
```golang
package main

import (
	"fmt"
	"net"

	"github.com/parinpan/protoevent"
)

func main() {
	agent, event := protoevent.CreateAgent("tcp", "0.0.0.0:8089")

	event.OnConnectionError(func(err error) {

	})

	event.OnConnectionAccepted(func(conn net.Conn) {
		fmt.Println("Accepting new connection: ", conn.RemoteAddr())
	})

	event.OnConnectionClosed(func(conn net.Conn) {

	})

	event.OnReceiveMessageError(func(conn net.Conn, err error) {

	})

	event.OnMessageReceived(func(conn net.Conn, message []byte) {
		fmt.Println("Received a message: ", string(message))
	})

	event.OnSendMessageError(func(conn net.Conn, message []byte, err error) {

	})

	event.OnMessageSent(func(conn net.Conn, message []byte) {
		fmt.Println("Sent a message: ", string(message))
	})

	err := agent.Run(func(conn net.Conn) error {
		_, err := conn.Write([]byte(`{"from": "AgentV1"}`))
		return err
	})

	if nil != err {
		panic(err)
	}
}
```
