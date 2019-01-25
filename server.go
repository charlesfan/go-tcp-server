package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/charlesfan/go-tcp-server/method"
	"github.com/charlesfan/go-tcp-server/service"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 3333, "The port to listen on; default is 3333.")

func main() {
	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)

	// Services init
	service.Init()
	// Server init
	e := method.Init()
	e.Run(src)
}
