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

type Connection struct {
	connectedAs connectionType
	connection  net.Conn
}

func NewConnection(connectedAs connectionType, conn net.Conn) *Connection {
	return &Connection{connectedAs, conn}
}

func (c *Connection) Read(b []byte) (n int, err error) {
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

func (c *Connection) Write(b []byte) (n int, err error) {
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

func (c *Connection) Close() error {
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

func (c *Connection) LocalAddr() net.Addr {
	return c.connection.LocalAddr()
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.connection.RemoteAddr()
}

func (c *Connection) SetDeadline(t time.Time) error {
	return c.connection.SetDeadline(t)
}

func (c *Connection) SetReadDeadline(t time.Time) error {
	return c.connection.SetReadDeadline(t)
}

func (c *Connection) SetWriteDeadline(t time.Time) error {
	return c.connection.SetWriteDeadline(t)
}
