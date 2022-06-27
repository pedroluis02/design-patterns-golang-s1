package abstract_factory

import "fmt"

type SocketServer interface {
	GetAddress() string
	ToString() string
}

//Production
type ProdSocketServer struct {
	Ip   string
	Port int
}

func (s *ProdSocketServer) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}

func (s *ProdSocketServer) ToString() string {
	return fmt.Sprintf("Socket{ip=%s port=%d}", s.Ip, s.Port)
}

//Testing
type TestSocketServer struct {
	Ip   string
	Port int
}

func (s *TestSocketServer) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}

func (s *TestSocketServer) ToString() string {
	return fmt.Sprintf("TestSocket{ip=%s port=%d}", s.Ip, s.Port)
}
