package service

var (
	HttpRequestService HttpRequestServicer
)

func Init() {
	// === Service ===
	// HttpRequest Service
	HttpRequestService = NewHttpRequestService()
	go HttpRequestService.Run()
}
