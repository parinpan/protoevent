package protoevent

import (
	"net"
)

func Listen(network, address string) (net.Listener, ServerEvent, error) {
	listener := new(listener)
	l, err := listener.Listen(network, address)
	return l, newServerEvent(), err
}

func CreateServant(network, address string) (Servant, ServerEvent, error) {
	l, event, err := Listen(network, address)
	return newServant(l), event, err
}

func CreateAgent(network, address string) (Agent, ClientEvent) {
	dialer := new(dialer)
	agent := newAgent(dialer.Dial, network, address)
	return agent, agent.GetEvent()
}
