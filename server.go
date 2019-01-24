package main

import (
	"fmt"

	"github.com/charlesfan/go-tcp-server/method"
	"github.com/charlesfan/go-tcp-server/service"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
)

func main() {
	fmt.Println("Starting server...")

	src := CONN_HOST + ":" + CONN_PORT

	// Services init
	service.Init()
	// Server init
	e := method.Init()
	e.Run(src)
}
