package protoevent

import (
	"net"
)

type Servant interface {
	Serve()
	SetMessageReadSize(size int)
}

type servantAgentImpl struct {
	listener        net.Listener
	messageReadSize int
}

func newServant(listener net.Listener) *servantAgentImpl {
	return &servantAgentImpl{listener, 1024}
}

func (sa *servantAgentImpl) SetMessageReadSize(size int) {
	sa.messageReadSize = size
}

func (sa *servantAgentImpl) Serve() {
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
