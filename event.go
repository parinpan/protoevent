package protoevent

import (
	"net"
)

type OnConnectionAcceptedExecFn func(conn net.Conn)
type OnConnectionClosedExecFn func(conn net.Conn)
type OnConnectionErrorExecFn func(err error)

type OnMessageReceivedExecFn func(conn net.Conn, message []byte, rawMessage []byte)
type OnReceiveMessageErrorExecFn func(conn net.Conn, err error)
type OnMessageSentExecFn func(conn net.Conn, message []byte)
type OnSendMessageErrorExecFn func(conn net.Conn, message []byte, err error)

var (
	// server events
	onServerConnectionAcceptedCallback  OnConnectionAcceptedExecFn  = func(conn net.Conn) {}
	onServerConnectionClosedCallback    OnConnectionClosedExecFn    = func(conn net.Conn) {}
	onServerConnectionErrorCallback     OnConnectionErrorExecFn     = func(err error) {}
	onServerMessageReceivedCallback     OnMessageReceivedExecFn     = func(conn net.Conn, message []byte, rawMessage []byte) {}
	onServerReceiveMessageErrorCallback OnReceiveMessageErrorExecFn = func(conn net.Conn, err error) {}
	onServerMessageSentCallback         OnMessageSentExecFn         = func(conn net.Conn, message []byte) {}
	onServerSendMessageErrorCallback    OnSendMessageErrorExecFn    = func(conn net.Conn, message []byte, err error) {}

	// client events
	onClientConnectionAcceptedCallback  OnConnectionAcceptedExecFn  = func(conn net.Conn) {}
	onClientConnectionClosedCallback    OnConnectionClosedExecFn    = func(conn net.Conn) {}
	onClientConnectionErrorCallback     OnConnectionErrorExecFn     = func(err error) {}
	onClientMessageReceivedCallback     OnMessageReceivedExecFn     = func(conn net.Conn, message []byte, rawMessage []byte) {}
	onClientReceiveMessageErrorCallback OnReceiveMessageErrorExecFn = func(conn net.Conn, err error) {}
	onClientMessageSentCallback         OnMessageSentExecFn         = func(conn net.Conn, message []byte) {}
	onClientSendMessageErrorCallback    OnSendMessageErrorExecFn    = func(conn net.Conn, message []byte, err error) {}
)

type EventBase interface {
	OnConnectionAccepted(fn OnConnectionAcceptedExecFn)
	OnConnectionClosed(fn OnConnectionClosedExecFn)
	OnMessageReceived(fn OnMessageReceivedExecFn)
	OnMessageSent(fn OnMessageSentExecFn)

	OnConnectionError(fn OnConnectionErrorExecFn)
	OnReceiveMessageError(fn OnReceiveMessageErrorExecFn)
	OnSendMessageError(fn OnSendMessageErrorExecFn)
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

func (se *serverEventImpl) OnConnectionError(fn OnConnectionErrorExecFn) {
	onServerConnectionErrorCallback = fn
}

func (se *serverEventImpl) OnSendMessageError(fn OnSendMessageErrorExecFn) {
	onServerSendMessageErrorCallback = fn
}

func (se *serverEventImpl) OnReceiveMessageError(fn OnReceiveMessageErrorExecFn) {
	onServerReceiveMessageErrorCallback = fn
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

func (ce *clientEventImpl) OnConnectionError(fn OnConnectionErrorExecFn) {
	onClientConnectionErrorCallback = fn
}

func (ce *clientEventImpl) OnSendMessageError(fn OnSendMessageErrorExecFn) {
	onClientSendMessageErrorCallback = fn
}

func (ce *clientEventImpl) OnReceiveMessageError(fn OnReceiveMessageErrorExecFn) {
	onClientReceiveMessageErrorCallback = fn
}
