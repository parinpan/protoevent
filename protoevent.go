package protoevent

func CreateServant(network, address string) (Servant, ServerEvent, error) {
	listener := newListener(network, address)

	l, err := listener.Listen()
	if nil != err {
		return nil, nil, err
	}

	return newServant(l), l.GetEvent(), err
}

func CreateAgent(network, address string) (Agent, ClientEvent) {
	dialer := newDialer(network, address)
	return newAgent(dialer), dialer.GetEvent()
}
