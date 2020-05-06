package protoevent

import (
	"net"
)

func Listen(network, address string) (net.Listener, ServerEvent, error) {
	listener := new(listener)
	l, err := listener.Listen(network, address)
	return l, newServerEvent(), err
}

func Dial(network, address string) (net.Conn, ClientEvent, error) {
	conn, err := net.Dial(network, address)

	if nil != err {
		defer onClientConnectionErrorCallback(err)
		return conn, newClientEvent(), err
	}

	newConnection := newConnection(clientConnection, conn)
	defer onClientConnectionAcceptedCallback(newConnection)

	return newConnection, newClientEvent(), nil
}

func CreateAgent(network, address string) (ServerAgent, ServerEvent, error) {
	l, event, err := Listen(network, address)
	return newServerAgent(l), event, err
}
