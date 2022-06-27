package abstract_factory

type ServerFactory interface {
	CreateSocketServer() SocketServer
	CreateApiRestFul() ApiRestFul
}
