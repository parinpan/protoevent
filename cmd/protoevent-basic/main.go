package main

import (
	"fmt"
	"net"

	"github.com/parinpan/protoevent"
)

func main() {
	listener, event, err := protoevent.Listen("tcp", "0.0.0.0:8089")

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

	for {
		conn, err := listener.Accept()

		if nil != err {
			panic(err)
		}

		go func() {
			data := make([]byte, 1024)
			conn.Read(data)
		}()
	}
}
