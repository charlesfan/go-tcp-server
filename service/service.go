package service

import (
	"github.com/charlesfan/go-tcp-server/config"
)

var (
	HttpRequestService HttpRequestServicer
	PixabayService     PixabayServicer
)

func Init() {
	// === Confgure ===
	cf := config.NewConfig()
	if e := cf.Init(); e == false {
		panic("Config init error")
	}
	// === Service ===
	// HttpRequest Service
	HttpRequestService = NewHttpRequestService()
	go HttpRequestService.Run()
	// Pixabay Service
	PixabayService = NewPixabayService(cf.Config.Pixabay.Key, cf.Config.Pixabay.URL, HttpRequestService)
}
