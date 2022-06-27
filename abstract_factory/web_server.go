package abstract_factory

import "fmt"

type ApiRestFul interface {
	GetAddress() string
	GetVersion() string
	ToString() string
}

//Production
type ProdApiRestFul struct {
	BaseUrl string
	Version int
}

func (a *ProdApiRestFul) GetAddress() string {
	return fmt.Sprintf("%s/%s", a.BaseUrl, a.GetVersion())
}

func (a *ProdApiRestFul) GetVersion() string {
	return fmt.Sprintf("v%d", a.Version)
}

func (s *ProdApiRestFul) ToString() string {
	return fmt.Sprintf("ApiRest{base=%s v=%d}", s.BaseUrl, s.Version)
}

//Testing
type TestApiRestFul struct {
	BaseUrl string
	Version int
}

func (a *TestApiRestFul) GetAddress() string {
	return fmt.Sprintf("%s/%s", a.BaseUrl, a.GetVersion())
}

func (a *TestApiRestFul) GetVersion() string {
	return fmt.Sprintf("v%d-dev", a.Version)
}

func (s *TestApiRestFul) ToString() string {
	return fmt.Sprintf("TestApiRest{base=%s v=%d}", s.BaseUrl, s.Version)
}
