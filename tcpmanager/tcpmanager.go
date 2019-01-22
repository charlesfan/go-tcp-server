package tcpmanager

import (
	"bufio"
	"container/list"
	"fmt"
	"net"
	"os"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	Router

	frontFunc  HandlerFunc
	handlerMap map[string]*list.List
	Listener   net.Listener
}

func New() *Engine {
	e := &Engine{
		handlerMap: make(map[string]*list.List),
	}
	e.Router.engine = e

	var initFunc HandlerFunc = func(c *Context) {
		fmt.Println("Got message:", c.message)
		if len(c.message) > 0 {
			c.elem.Value.(HandlerFunc)(c)
			return
		}
	}

	e.frontFunc = initFunc

	return e
}

func NewListener(host string) (*Engine, error) {
	e := &Engine{
		handlerMap: make(map[string]*list.List),
	}
	e.Router.engine = e

	listener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return nil, err
	}
	e.Listener = listener

	var initFunc HandlerFunc = func(c *Context) {
		fmt.Println("Got message:", c.message)
		if len(c.message) > 0 {
			c.elem.Value.(HandlerFunc)(c)
			return
		}
	}

	e.frontFunc = initFunc

	return e, nil
}

var _ IRouter = &Engine{}

func (engine *Engine) Run(host string) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	engine.Listener = listener
	fmt.Printf("Listening on %s\n", host)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go engine.handleConnection(conn)
	}
}

func (engine *Engine) Stop() {
	engine.Listener.Close()
}

func (engin *Engine) GetHandlers() map[string]*list.List {
	return engin.handlerMap
}

func (engine *Engine) handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	scanner := bufio.NewScanner(conn)

	c := &Context{
		Conn:       conn,
		Scanner:    scanner,
		MethodName: "root",
	}

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		engine.handleContext(c)
	}
	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func (engine *Engine) handleContext(c *Context) {
	message := c.Scanner.Text()

	if len(message) <= 0 {
		return
	}

	switch message {
	case "quit":
		fmt.Println("Quitting.")
		c.Conn.Write([]byte("Disconnection!!!\n"))
		c.Conn.Close()
		return
	case "pwd":
		c.Conn.Write([]byte(c.MethodName + "\n"))
		return
	}

	if len(message) > 0 && c.MethodName != "root" {
		c.message = message
		c.elem.Value.(HandlerFunc)(c)
		return
	}

	a := strings.SplitN(message, " ", 2)

	if engine.handlerMap[a[0]] == nil {
		c.Conn.Write([]byte("Unrecognized command.\n"))
		return
	}

	if len(a) > 1 {
		c.message = a[1]
	}

	len := (*engine.handlerMap[a[0]]).Len()
	switch {
	case len <= 0:
		c.Conn.Write([]byte("Without handler command.\n"))
		return
	case len == 1:
		engine.handlerMap[a[0]].Front().Value.(HandlerFunc)(c)
		return
	}

	e := engine.handlerMap[a[0]].Front()
	c.MethodName = a[0]
	c.elem = e.Next()
	e.Value.(HandlerFunc)(c)
}
