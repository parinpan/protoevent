package protoevent

import (
	"net"
)

type Servant interface {
	Serve()
	SetDefaultReadSize(size int)
}

type servantAgentImpl struct {
	listener        net.Listener
	defaultReadSize int
}

func newServant(listener net.Listener) *servantAgentImpl {
	return &servantAgentImpl{listener, 1024}
}

func (sa *servantAgentImpl) SetDefaultReadSize(size int) {
	sa.defaultReadSize = size
}

func (sa *servantAgentImpl) Serve() {
	for {
		conn, err := sa.listener.Accept()

		if nil != err {
			continue
		}

		go func() {
			var buffer = make([]byte, sa.defaultReadSize)

			for {
				defer conn.Close()

				if _, err := conn.Read(buffer); nil != err {
					break
				}
			}
		}()
	}
}
