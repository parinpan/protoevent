package protoevent

import (
	"net"
)

type Servant interface {
	Serve()
	SetDefaultReadSize(size int)
}

type servantImpl struct {
	listener        net.Listener
	defaultReadSize int
}

func newServant(listener net.Listener) *servantImpl {
	return &servantImpl{listener, 1024}
}

func (s *servantImpl) SetDefaultReadSize(size int) {
	s.defaultReadSize = size
}

func (s *servantImpl) Serve() {
	for {
		conn, err := s.listener.Accept()

		if nil != err {
			continue
		}

		go func() {
			var buffer = make([]byte, s.defaultReadSize)

			for {
				defer conn.Close()

				if _, err := conn.Read(buffer); nil != err {
					break
				}
			}
		}()
	}
}
