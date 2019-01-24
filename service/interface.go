package service

type HttpRequestServicer interface {
	Put(*Content)
	Run()
	Stop()
}
