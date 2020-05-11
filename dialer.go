package protoevent

import (
	"net"
)

type dialer struct {
	network string
	address string
	event   *networkEvent
}

func newDialer(network, address string) *dialer {
	return &dialer{
		network: network,
		address: address,
		event:   newNetworkEvent(clientConnection, address),
	}
}

func (d *dialer) Dial() (*connection, error) {
	conn, err := net.Dial(d.network, d.address)

	if nil != err {
		d.event.GetCallbackStorage().OnConnectionError(err)
		return nil, err
	}

	newConnection := newConnection(clientConnection, conn, d.event)
	d.event.GetCallbackStorage().OnConnectionAccepted(newConnection)

	return newConnection, err
}

func (d *dialer) GetEvent() ClientEvent {
	return d.event
}
