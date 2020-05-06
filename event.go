package protoevent

import (
	"net"
)

type OnConnectionAcceptedExecFn func(conn net.Conn)
type OnConnectionClosedExecFn func(conn net.Conn)
type OnMessageReceivedExecFn func(conn net.Conn, message []byte)
type OnMessageSentExecFn func(conn net.Conn, message []byte)

var (
	// server events
	onServerConnectionAcceptedCallback OnConnectionAcceptedExecFn = func(conn net.Conn) {}
	onServerConnectionClosedCallback   OnConnectionClosedExecFn   = func(conn net.Conn) {}
	onServerMessageReceivedCallback    OnMessageReceivedExecFn    = func(conn net.Conn, message []byte) {}
	onServerMessageSentCallback        OnMessageSentExecFn        = func(conn net.Conn, message []byte) {}

	// client events
	onClientConnectionAcceptedCallback OnConnectionAcceptedExecFn = func(conn net.Conn) {}
	onClientConnectionClosedCallback   OnConnectionClosedExecFn   = func(conn net.Conn) {}
	onClientMessageReceivedCallback    OnMessageReceivedExecFn    = func(conn net.Conn, message []byte) {}
	onClientMessageSentCallback        OnMessageSentExecFn        = func(conn net.Conn, message []byte) {}
)

type EventBase interface {
	OnConnectionAccepted(fn OnConnectionAcceptedExecFn)
	OnConnectionClosed(fn OnConnectionClosedExecFn)
	OnMessageReceived(fn OnMessageReceivedExecFn)
	OnMessageSent(fn OnMessageSentExecFn)
}

type ServerEvent interface {
	EventBase
}

type ClientEvent interface {
	EventBase
}

type serverEventImpl struct {
}

func newServerEvent() *serverEventImpl {
	return new(serverEventImpl)
}

func (se *serverEventImpl) OnConnectionAccepted(fn OnConnectionAcceptedExecFn) {
	onServerConnectionAcceptedCallback = fn
}

func (se *serverEventImpl) OnConnectionClosed(fn OnConnectionClosedExecFn) {
	onServerConnectionClosedCallback = fn
}

func (se *serverEventImpl) OnMessageReceived(fn OnMessageReceivedExecFn) {
	onServerMessageReceivedCallback = fn
}

func (se *serverEventImpl) OnMessageSent(fn OnMessageSentExecFn) {
	onServerMessageSentCallback = fn
}

type clientEventImpl struct {
}

func newClientEvent() *clientEventImpl {
	return new(clientEventImpl)
}

func (ce *clientEventImpl) OnConnectionAccepted(fn OnConnectionAcceptedExecFn) {
	onClientConnectionAcceptedCallback = fn
}

func (ce *clientEventImpl) OnConnectionClosed(fn OnConnectionClosedExecFn) {
	onClientConnectionClosedCallback = fn
}

func (ce *clientEventImpl) OnMessageReceived(fn OnMessageReceivedExecFn) {
	onClientMessageReceivedCallback = fn
}

func (ce *clientEventImpl) OnMessageSent(fn OnMessageSentExecFn) {
	onClientMessageSentCallback = fn
}
