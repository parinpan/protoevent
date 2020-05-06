package main

import (
	"fmt"
	"net"

	"github.com/parinpan/protoevent"
)

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
	})

	event.OnSendMessageError(func(conn net.Conn, message []byte, err error) {

	})

	event.OnMessageSent(func(conn net.Conn, message []byte) {

	})

	servant.Serve()
}
