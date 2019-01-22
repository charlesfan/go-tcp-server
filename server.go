package main

import (
	"fmt"

	"github.com/charlesfan/go-tcp-server/method"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
)

func main() {
	fmt.Println("Starting server...")

	src := CONN_HOST + ":" + CONN_PORT

	e := method.Init()
	e.Run(src)
}
