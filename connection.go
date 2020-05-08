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

	go func() {
		copiedBytes := b

		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageReceivedCallback(c, copiedBytes[:n])
			case clientConnection:
				onClientMessageReceivedCallback(c, copiedBytes[:n])
			}
		} else {
			switch c.connectedAs {
			case serverConnection:
				onServerReceiveMessageErrorCallback(c, err)
			case clientConnection:
				onClientReceiveMessageErrorCallback(c, err)
			}
		}
	}()

	return n, err
}

func (c *connection) Write(b []byte) (n int, err error) {
	n, err = c.connection.Write(b)

	go func() {
		copiedBytes := b

		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageSentCallback(c, copiedBytes)
			case clientConnection:
				onClientMessageSentCallback(c, copiedBytes)
			}
		} else {
			switch c.connectedAs {
			case serverConnection:
				onServerSendMessageErrorCallback(c, copiedBytes, err)
			case clientConnection:
				onClientSendMessageErrorCallback(c, copiedBytes, err)
			}
		}
	}()

	return n, err
}

func (c *connection) Close() error {
	err := c.connection.Close()

	go func() {
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
