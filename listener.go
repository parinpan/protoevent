package protoevent

import (
	"net"
)

type Listener struct {
	listener net.Listener
}

func (l *Listener) Listen(network, address string) (*Listener, error) {
	var err error
	l.listener, err = net.Listen(network, address)
	return l, err
}

func (l *Listener) Accept() (net.Conn, error) {
	conn, err := l.listener.Accept()

	if nil != err {
		return conn, err
	}

	newConnection := NewConnection(serverConnection, conn)
	onServerConnectionAcceptedCallback(newConnection)

	return newConnection, err
}

func (l *Listener) Close() error {
	return l.listener.Close()
}

func (l *Listener) Addr() net.Addr {
	return l.listener.Addr()
}
