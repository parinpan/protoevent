package protoevent

import (
	"net"
)

type dialer struct {
	conn *connection
}

func (d *dialer) Dial(network, address string) (*connection, error) {
	conn, err := net.Dial(network, address)
	wrappedConn := newConnection(clientConnection, conn)
	
	if nil != err {
		onClientConnectionErrorCallback(err)
		return wrappedConn, err
	}
	
	return wrappedConn, err
}
