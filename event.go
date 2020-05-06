package protoevent

import (
	"net"
)

type OnConnectionAcceptedExecFn func(conn net.Conn)
type OnConnectionClosedExecFn func(conn net.Conn)
type OnMessageReceivedExecFn func(conn net.Conn, message []byte)
type OnMessageSentExecFn func(conn net.Conn, message []byte)

type OnConnectionAcceptedRegFn func(execFn OnConnectionAcceptedExecFn)
type OnConnectionClosedRegFn func(execFn OnConnectionClosedExecFn)
type OnMessageReceivedRegFn func(execFn OnMessageReceivedExecFn)
type OnMessageSentRegFn func(execFn OnMessageSentExecFn)

var (
	// server events
	onServerConnectionAcceptedCallback OnConnectionAcceptedExecFn
	onServerConnectionClosedCallback   OnConnectionClosedExecFn
	onServerMessageReceivedCallback    OnMessageReceivedExecFn
	onServerMessageSentCallback        OnMessageSentExecFn

	// client events
	onClientConnectionAcceptedCallback OnConnectionAcceptedExecFn
	onClientConnectionClosedCallback   OnConnectionClosedExecFn
	onClientMessageReceivedCallback    OnMessageReceivedExecFn
	onClientMessageSentCallback        OnMessageSentExecFn
)

func GetClientEventRegistrars() (OnConnectionAcceptedRegFn, OnConnectionClosedRegFn, OnMessageReceivedRegFn, OnMessageSentRegFn) {
	return setOnClientConnectionAcceptedExecFn, setOnClientConnectionClosedExecFn,
		setOnClientMessageReceivedExecFn, setOnClientMessageSentExecFn
}

func GetServerEventRegistrars() (OnConnectionAcceptedRegFn, OnConnectionClosedRegFn, OnMessageReceivedRegFn, OnMessageSentRegFn) {
	return setOnServerConnectionAcceptedExecFn, setOnServerConnectionClosedExecFn,
		setOnServerMessageReceivedExecFn, setOnServerMessageSentExecFn
}

func setOnServerConnectionAcceptedExecFn(execFn OnConnectionAcceptedExecFn) {
	onServerConnectionAcceptedCallback = execFn
}

func setOnServerConnectionClosedExecFn(execFn OnConnectionClosedExecFn) {
	onServerConnectionClosedCallback = execFn
}

func setOnServerMessageReceivedExecFn(execFn OnMessageReceivedExecFn) {
	onServerMessageReceivedCallback = execFn
}

func setOnServerMessageSentExecFn(execFn OnMessageSentExecFn) {
	onServerMessageSentCallback = execFn
}

func setOnClientConnectionAcceptedExecFn(execFn OnConnectionAcceptedExecFn) {
	onClientConnectionAcceptedCallback = execFn
}

func setOnClientConnectionClosedExecFn(execFn OnConnectionClosedExecFn) {
	onClientConnectionClosedCallback = execFn
}

func setOnClientMessageReceivedExecFn(execFn OnMessageReceivedExecFn) {
	onClientMessageReceivedCallback = execFn
}

func setOnClientMessageSentExecFn(execFn OnMessageSentExecFn) {
	onClientMessageSentCallback = execFn
}
