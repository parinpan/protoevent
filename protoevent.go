package protoevent

func CreateServant(network, address string) (Servant, ServerEvent, error) {
	var err error
	var listener = newListener(network, address)

	if err = listener.Listen(); nil != err {
		return nil, nil, err
	}

	return newServant(listener), listener.GetEvent(), err
}

func CreateAgent(network, address string) (Agent, ClientEvent) {
	dialer := newDialer(network, address)
	return newAgent(dialer), dialer.GetEvent()
}
