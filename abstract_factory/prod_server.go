package abstract_factory

type ProductionServer struct {
}

func (s *ProductionServer) CreateSocketServer() SocketServer {
	return &ProdSocketServer{
		Ip:   "10.9.8.1",
		Port: 7000,
	}
}

func (s *ProductionServer) CreateApiRestFul() ApiRestFul {
	return &ProdApiRestFul{
		BaseUrl: "http://prodmyapi.com",
		Version: 1,
	}
}
