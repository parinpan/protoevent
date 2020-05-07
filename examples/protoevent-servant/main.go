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
