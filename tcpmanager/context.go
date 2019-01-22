package tcpmanager

import (
	"bufio"
	"container/list"
	"net"
)

type Context struct {
	Conn       net.Conn
	MethodName string
	Scanner    *bufio.Scanner
	message    string
	elem       *list.Element
}

func (c *Context) GetMessage() string {
	return c.message
}

func (c *Context) Clear() {
	c.message = ""
	c.MethodName = "root"
}
