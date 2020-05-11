package protoevent

import (
	"net"
	"time"
)

type connectionType string

const (
	serverConnection connectionType = "serverConnection"
	clientConnection connectionType = "clientConnection"
)

type connection struct {
	connectedAs connectionType
	connection  net.Conn
	event       *networkEvent
}

func newConnection(connectedAs connectionType, conn net.Conn, event *networkEvent) *connection {
	return &connection{
		connectedAs: connectedAs,
		connection:  conn,
		event:       event,
	}
}

func (c *connection) Read(b []byte) (n int, err error) {
	n, err = c.connection.Read(b)

	go func() {
		copiedBytes := b

		if nil == err {
			c.event.GetCallbackStorage().OnMessageReceived(c, copiedBytes[:n], copiedBytes)
		} else {
			c.event.GetCallbackStorage().OnReceiveMessageError(c, err)
		}
	}()

	return n, err
}

func (c *connection) Write(b []byte) (n int, err error) {
	n, err = c.connection.Write(b)

	go func() {
		copiedBytes := b

		if nil == err {
			c.event.GetCallbackStorage().OnMessageSent(c, copiedBytes)
		} else {
			c.event.GetCallbackStorage().OnSendMessageError(c, copiedBytes, err)
		}
	}()

	return n, err
}

func (c *connection) Close() error {
	err := c.connection.Close()

	go func() {
		if nil == err {
			c.event.GetCallbackStorage().OnConnectionClosed(c)
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
