package protoevent

import (
	"net"
)

type listener struct {
	listener net.Listener
}

func (l *listener) Listen(network, address string) (*listener, error) {
	var err error
	l.listener, err = net.Listen(network, address)
	return l, err
}

func (l *listener) Accept() (net.Conn, error) {
	conn, err := l.listener.Accept()

	if nil != err {
		defer onServerConnectionErrorCallback(err)
		return conn, err
	}

	newConnection := newConnection(serverConnection, conn)
	defer onServerConnectionAcceptedCallback(newConnection)

	return newConnection, err
}

func (l *listener) Close() error {
	return l.listener.Close()
}

func (l *listener) Addr() net.Addr {
	return l.listener.Addr()
}
