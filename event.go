package protoevent

import (
	"fmt"
	"net"
)

var (
	eventCallbackStorageMap = make(map[string]*eventCallbackStorage)
)

type OnConnectionAcceptedExecFn func(conn net.Conn)
type OnConnectionClosedExecFn func(conn net.Conn)
type OnConnectionErrorExecFn func(err error)

type OnReceiveMessageErrorExecFn func(conn net.Conn, err error)
type OnMessageReceivedExecFn func(conn net.Conn, message []byte, rawMessage []byte)

type OnSendMessageErrorExecFn func(conn net.Conn, message []byte, err error)
type OnMessageSentExecFn func(conn net.Conn, message []byte)

type EventBase interface {
	OnConnectionError(fn OnConnectionErrorExecFn)
	OnConnectionAccepted(fn OnConnectionAcceptedExecFn)
	OnConnectionClosed(fn OnConnectionClosedExecFn)

	OnMessageReceived(fn OnMessageReceivedExecFn)
	OnReceiveMessageError(fn OnReceiveMessageErrorExecFn)

	OnMessageSent(fn OnMessageSentExecFn)
	OnSendMessageError(fn OnSendMessageErrorExecFn)
}

type ServerEvent interface {
	EventBase
	GetCallbackStorage() *eventCallbackStorage
}

type ClientEvent interface {
	EventBase
	GetCallbackStorage() *eventCallbackStorage
}

type eventCallbackStorage struct {
	OnConnectionError    OnConnectionErrorExecFn
	OnConnectionAccepted OnConnectionAcceptedExecFn
	OnConnectionClosed   OnConnectionClosedExecFn

	OnReceiveMessageError OnReceiveMessageErrorExecFn
	OnMessageReceived     OnMessageReceivedExecFn

	OnSendMessageError OnSendMessageErrorExecFn
	OnMessageSent      OnMessageSentExecFn
}

func newEventStore(connectedAs connectionType, network string) *eventCallbackStorage {
	hashKey := fmt.Sprint(connectedAs, network)

	eventCallbackStorageMap[hashKey] = &eventCallbackStorage{
		OnConnectionError:     func(err error) {},
		OnConnectionAccepted:  func(conn net.Conn) {},
		OnConnectionClosed:    func(conn net.Conn) {},
		OnReceiveMessageError: func(conn net.Conn, err error) {},
		OnMessageReceived:     func(conn net.Conn, message []byte, rawMessage []byte) {},
		OnSendMessageError:    func(conn net.Conn, message []byte, err error) {},
		OnMessageSent:         func(conn net.Conn, message []byte) {},
	}

	fmt.Println(hashKey, eventCallbackStorageMap[hashKey])

	return eventCallbackStorageMap[hashKey]
}

type networkEvent struct {
	eventCallbackStorage *eventCallbackStorage
}

func newNetworkEvent(connectedAs connectionType, address string) *networkEvent {
	return &networkEvent{
		eventCallbackStorage: newEventStore(connectedAs, address),
	}
}

func (ne *networkEvent) OnConnectionError(fn OnConnectionErrorExecFn) {
	ne.eventCallbackStorage.OnConnectionError = fn
}

func (ne *networkEvent) OnConnectionAccepted(fn OnConnectionAcceptedExecFn) {
	ne.eventCallbackStorage.OnConnectionAccepted = fn
}

func (ne *networkEvent) OnConnectionClosed(fn OnConnectionClosedExecFn) {
	ne.eventCallbackStorage.OnConnectionClosed = fn
}

func (ne *networkEvent) OnReceiveMessageError(fn OnReceiveMessageErrorExecFn) {
	ne.eventCallbackStorage.OnReceiveMessageError = fn
}

func (ne *networkEvent) OnMessageReceived(fn OnMessageReceivedExecFn) {
	ne.eventCallbackStorage.OnMessageReceived = fn
}

func (ne *networkEvent) OnSendMessageError(fn OnSendMessageErrorExecFn) {
	ne.eventCallbackStorage.OnSendMessageError = fn
}

func (ne *networkEvent) OnMessageSent(fn OnMessageSentExecFn) {
	ne.eventCallbackStorage.OnMessageSent = fn
}

func (ne *networkEvent) GetCallbackStorage() *eventCallbackStorage {
	return ne.eventCallbackStorage
}
