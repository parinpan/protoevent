package protoevent

import (
	"net"
)

type ServerAgent interface {
	Run()
	SetMessageReadSize(size int)
}

type serverAgentImpl struct {
	listener        net.Listener
	messageReadSize int
}

func newServerAgent(listener net.Listener) *serverAgentImpl {
	return &serverAgentImpl{listener, 1024}
}

func (sa *serverAgentImpl) SetMessageReadSize(size int) {
	sa.messageReadSize = size
}

func (sa *serverAgentImpl) Run() {
	for {
		conn, err := sa.listener.Accept()

		if nil != err {
			continue
		}

		go func() {
			conn.Read(make([]byte, sa.messageReadSize))
		}()
	}
}
