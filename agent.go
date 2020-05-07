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
	connection net.Conn
	defaultReadSize int
}

func newAgent(conn net.Conn) *agentImpl {
	return &agentImpl{conn, 1024}
}

func (a *agentImpl) SetDefaultReadSize(size int) {
	a.defaultReadSize = size
}

func (a *agentImpl) Run(fn triggerFn) error {
	err := fn(a.connection)
	
	if nil != err {
		return err
	}
	
	for {
		message := make([]byte, a.defaultReadSize)
		_, err = a.connection.Read(message)
		
		if nil != err {
			break
		}
	}
	
	return err
}
