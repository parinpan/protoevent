package protoevent

import (
	"net"
)

type triggerFn func(conn net.Conn) error
type dialerFn func(network, address string) (*connection, error)

type Agent interface {
	Run(fn triggerFn) error
	GetEvent() ClientEvent
	SetDefaultReadSize(size int)
}

type agentImpl struct {
	dialer     dialerFn
	connection net.Conn
	event      ClientEvent

	defaultReadSize int
	network         string
	address         string
}

func newAgent(dialer dialerFn, network, address string) *agentImpl {
	return &agentImpl{
		network:         network,
		address:         address,
		dialer:          dialer,
		defaultReadSize: 1024,
		event:           newClientEvent(),
	}
}

func (a *agentImpl) SetDefaultReadSize(size int) {
	a.defaultReadSize = size
}

func (a *agentImpl) GetEvent() ClientEvent {
	return a.event
}

func (a *agentImpl) Run(fn triggerFn) error {
	var err error
	a.connection, err = a.dialer(a.network, a.address)

	if nil != err {
		return err
	}

	defer a.connection.Close()

	if err = fn(a.connection); nil != err {
		return err
	}

	message := make([]byte, a.defaultReadSize)

	for {
		_, err = a.connection.Read(message)

		if nil != err {
			break
		}
	}

	return err
}
