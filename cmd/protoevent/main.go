package main

import (
	"fmt"
	"net"

	"github.com/parinpan/protoevent"
)

func main() {
	registerProtoEvents()
	listener, err := protoevent.Listen("tcp", "0.0.0.0:8089")

	if nil != err {
		panic(err)
	}

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

func registerProtoEvents() {
	onConnectionAccepted, onConnectionClosed, onMessageReceived, onMessageSent := protoevent.GetServerEventRegistrars()

	onConnectionAccepted(func(conn net.Conn) {
		fmt.Println("Accepting new connection: ", conn.RemoteAddr())
	})

	onConnectionClosed(func(conn net.Conn) {

	})

	onMessageReceived(func(conn net.Conn, message []byte) {
		fmt.Println("Received a message: ", string(message))
	})

	onMessageSent(func(conn net.Conn, message []byte) {

	})
}
