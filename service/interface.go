package service

import (
	"net"
)

type HttpRequestServicer interface {
	Put(*Content)
	Run()
	Stop()
}

type PixabayServicer interface {
	GetPhotos(net.Conn, string)
	GetVideos(net.Conn, string)
	Alive() bool
}
