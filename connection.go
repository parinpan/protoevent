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
		copiedConn, copiedBytes := *c, b

		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageReceivedCallback(&copiedConn, copiedBytes[:n])
			case clientConnection:
				onClientMessageReceivedCallback(&copiedConn, copiedBytes[:n])
			}
		} else {
			switch c.connectedAs {
			case serverConnection:
				onServerReceiveMessageErrorCallback(&copiedConn, err)
			case clientConnection:
				onClientReceiveMessageErrorCallback(&copiedConn, err)
			}
		}
	}()

	return n, err
}

func (c *connection) Write(b []byte) (n int, err error) {
	n, err = c.connection.Write(b)

	go func() {
		copiedConn, copiedBytes := *c, b

		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerMessageSentCallback(&copiedConn, copiedBytes)
			case clientConnection:
				onClientMessageSentCallback(&copiedConn, copiedBytes)
			}
		} else {
			switch c.connectedAs {
			case serverConnection:
				onServerSendMessageErrorCallback(&copiedConn, copiedBytes, err)
			case clientConnection:
				onClientSendMessageErrorCallback(&copiedConn, copiedBytes, err)
			}
		}
	}()

	return n, err
}

func (c *connection) Close() error {
	err := c.connection.Close()

	go func() {
		copiedConn := *c

		if nil == err {
			switch c.connectedAs {
			case serverConnection:
				onServerConnectionClosedCallback(&copiedConn)
			case clientConnection:
				onClientConnectionClosedCallback(&copiedConn)
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
