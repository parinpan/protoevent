package protoevent

import (
	"net"
)

type dialer struct {
	conn *connection
}

func (d *dialer) Dial(network, address string) (*connection, error) {
	conn, err := net.Dial(network, address)

	if nil != err {
		onClientConnectionErrorCallback(err)
		return nil, err
	}

	wrappedConn := newConnection(clientConnection, conn)
	onClientConnectionAcceptedCallback(wrappedConn)

	return wrappedConn, err
}
