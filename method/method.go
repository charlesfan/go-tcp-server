package method

import (
	"github.com/charlesfan/go-tcp-server/tcpmanager"
)

func Init() *tcpmanager.Engine {
	e := tcpmanager.New()

	e.NewMethod("photo", PhotoHandler, false)
	e.NewMethod("video", VideoHandler, false)
	e.NewMethod("pixabay", PixabayAliveHandler, true)

	return e
}
