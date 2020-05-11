package protoevent

func CreateServant(network, address string) (Servant, ServerEvent, error) {
	listener := newListener(network, address)

	if err := listener.Listen(); nil != err {
		return nil, nil, err
	}

	return newServant(listener), listener.GetEvent(), nil
}

func CreateAgent(network, address string) (Agent, ClientEvent) {
	dialer := newDialer(network, address)
	return newAgent(dialer), dialer.GetEvent()
}
