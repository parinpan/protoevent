package protoevent

import (
	"net"
)

type triggerFn func(conn net.Conn) error

type Agent interface {
	Run(fn triggerFn) error
	SetDefaultReadSize(size int)
}

type agentImpl struct {
	connection      net.Conn
	dialer          *dialer
	defaultReadSize int
}

func newAgent(dialer *dialer) *agentImpl {
	return &agentImpl{
		dialer:          dialer,
		defaultReadSize: 1024,
	}
}

func (a *agentImpl) SetDefaultReadSize(size int) {
	a.defaultReadSize = size
}

func (a *agentImpl) Run(fn triggerFn) error {
	var err error
	a.connection, err = a.dialer.Dial()

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
