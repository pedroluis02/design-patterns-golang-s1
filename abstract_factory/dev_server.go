package abstract_factory

type TestingServer struct {
}

func (s *TestingServer) CreateSocketServer() SocketServer {
	return &TestSocketServer{
		Ip:   "localhost",
		Port: 5000,
	}
}

func (s *TestingServer) CreateApiRestFul() ApiRestFul {
	return &TestApiRestFul{
		BaseUrl: "http://localhost",
		Version: 1,
	}
}
