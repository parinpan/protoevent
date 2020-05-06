package protoevent

import (
	"net"
)

func Listen(network, address string) (net.Listener, error) {
	listener := new(listener)
	return listener.Listen(network, address)
}

func Dial(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)

	if nil != err {
		return conn, err
	}

	newConnection := newConnection(clientConnection, conn)
	onClientConnectionAcceptedCallback(newConnection)

	return newConnection, nil
}
