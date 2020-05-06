package protoevent

import (
	"net"
	"time"
)

type connectionType string

const (
	serverConnection = "serverConnection"
	clientConnection = "clientConnection"
)

type connection struct {
	connectedAs connectionType
	connection  net.Conn
}

func newConnection(connectedAs connectionType, conn net.Conn) *connection {
	return &connection{connectedAs, conn}
}

func (c *connection) Read(b []byte) (n int, err error) {
	n, err = c.connection.Read(b)

	defer func() {
		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageReceivedCallback(c, b)
			case clientConnection:
				onClientMessageReceivedCallback(c, b)
			}
		}
	}()

	return n, err
}

func (c *connection) Write(b []byte) (n int, err error) {
	n, err = c.connection.Write(b)

	defer func() {
		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageSentCallback(c, b)
			case clientConnection:
				onClientMessageSentCallback(c, b)
			}
		}
	}()

	return n, err
}

func (c *connection) Close() error {
	err := c.connection.Close()

	defer func() {
		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerConnectionClosedCallback(c)
			case clientConnection:
				onClientConnectionClosedCallback(c)
			}
		}
	}()

	return err
}

func (c *connection) LocalAddr() net.Addr {
	return c.connection.LocalAddr()
}

func (c *connection) RemoteAddr() net.Addr {
	return c.connection.RemoteAddr()
}

func (c *connection) SetDeadline(t time.Time) error {
	return c.connection.SetDeadline(t)
}

func (c *connection) SetReadDeadline(t time.Time) error {
	return c.connection.SetReadDeadline(t)
}

func (c *connection) SetWriteDeadline(t time.Time) error {
	return c.connection.SetWriteDeadline(t)
}
