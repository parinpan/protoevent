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
	dialer := new(dialer)
	conn, err := dialer.Dial(network, address)
	return conn, newClientEvent(), err
}

func CreateServant(network, address string) (Servant, ServerEvent, error) {
	l, event, err := Listen(network, address)
	return newServant(l), event, err
}

func CreateAgent(network, address string) (Agent, ClientEvent, error) {
	conn, event, err := Dial(network, address)
	return newAgent(conn), event, err
}
