package protoevent

import (
	"net"
)

type listener struct {
	network string
	address string

	listener net.Listener
	event    *networkEvent
}

func newListener(network, address string) *listener {
	return &listener{
		address: address,
		network: network,
		event:   newNetworkEvent(),
	}
}

func (l *listener) Listen() (*listener, error) {
	var err error
	l.listener, err = net.Listen(l.network, l.address)
	return l, err
}

func (l *listener) Accept() (net.Conn, error) {
	conn, err := l.listener.Accept()

	if nil != err {
		l.event.GetCallbackStorage().OnConnectionError(err)
		return nil, err
	}

	newConnection := newConnection(serverConnection, conn, l.event)
	l.event.GetCallbackStorage().OnConnectionAccepted(newConnection)

	return newConnection, err
}

func (l *listener) Close() error {
	return l.listener.Close()
}

func (l *listener) Addr() net.Addr {
	return l.listener.Addr()
}

func (l *listener) GetEvent() ServerEvent {
	return l.event
}
