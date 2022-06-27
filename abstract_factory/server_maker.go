package abstract_factory

const (
	Testing = iota
	Production
)

type ServerFactoryMaker struct {
}

func (f *ServerFactoryMaker) GetFactory(env int) ServerFactory {
	if env == Testing {
		return &TestingServer{}
	} else if env == Production {
		return &ProductionServer{}
	}

	return nil
}
