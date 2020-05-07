<p align="center"> 
	<img src="https://user-images.githubusercontent.com/14908455/81298347-dcb05e80-909e-11ea-8aa3-44ecf2c46af9.png"/>
</p>

ProtoEvent is an event-based TCP/UDP connection handling in Golang. It's simple and not exploiting too much the basic functionality of making a new protocol connection. Basically, ProtoEvent reimplements `net.Listener` and `net.Conn`  interface to have extended ability of capturing various events happened in network communication.

# Background
It's painful enough when you have to look after your simple network application codebase which is not well-organized ends up with producing more spaghetti code. At least, we're now able to organize the application logics based on a certain event.

# Installation
```
go get -u https://github.com/parinpan/protoevent
```

# Demo
![ezgif com-video-to-gif](https://user-images.githubusercontent.com/14908455/81297569-a0303300-909d-11ea-90f4-935d1d16925b.gif)

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

	servant.SetDefaultReadSize(4096) // set default read size per chunk in bytes
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
	agent.SetDefaultReadSize(4096) // set default read size per chunk in bytes

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

	// trigger a message to get connected with the server at first
	err := agent.Run(func(conn net.Conn) error {
		_, err := conn.Write([]byte(`{"from": "AgentV1"}`))
		return err
	})

	if nil != err {
		panic(err)
	}
}
```
# Contact Me 
If you have any inquiries. You can catch me up on:
- [Linkedin](https://linkedin.com/in/fachrinfan)
- [Twitter](https://twitter.com/fachrinfan)
- [Instagram](https://instagram.com/fachrinfan)
- [Email](mailto:fachrin.nasution@gmail.com)
