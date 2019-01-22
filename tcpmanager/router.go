package tcpmanager

import (
	"container/list"
)

type IRouter interface {
	NewMethod(string, HandlerFunc, bool) IRouter
}

type Router struct {
	engine *Engine
}

var _ IRouter = &Router{}

func (r *Router) NewMethod(method string, fn HandlerFunc, root bool) IRouter {
	if len(method) <= 0 {
		panic("Method name is Null")
	}
	r.newHandler(method, fn, root)
	return r.engine
}

func (r *Router) newHandler(method string, fn HandlerFunc, root bool) {
	if r.engine.handlerMap[method] != nil {
		panic("Method is duplicate: " + method)
	}
	l := list.New()
	if root == false {
		first := l.PushFront(r.engine.frontFunc)
		l.InsertAfter(fn, first)
	} else {
		l.PushFront(fn)
	}
	r.engine.handlerMap[method] = l
}
