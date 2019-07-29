package service

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/charlesfan/go-tcp-server/tools"
)

type Content struct {
	Conn   net.Conn
	Method string
	Url    string
}

type httpRequestService struct {
	queue   chan *Content
	stop    chan bool
	limiter *tools.Limiter
}

var _ HttpRequestServicer = &httpRequestService{}

func (r *httpRequestService) Put(c *Content) {
	r.queue <- c
}

func (r *httpRequestService) Run() {
	for {
		if r.limiter.Limit() {
			select {
			case c := <-r.queue:
				r.contentHandler(c)
			case <-r.stop:
				return
			}
		}
	}
}

func (r *httpRequestService) Stop() {
	r.stop <- true
}

func (r *httpRequestService) contentHandler(c *Content) {
	client := &http.Client{}

	req, err := http.NewRequest(c.Method, c.Url, nil)
	if err != nil {
		c.Conn.Write([]byte(err.Error()))
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.Conn.Write([]byte(err.Error()))
		return
	}

	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Conn.Write([]byte(err.Error()))
		return
	}
	c.Conn.Write([]byte(s))
	c.Conn.Write([]byte("\n"))
}

func NewHttpRequestService() HttpRequestServicer {
	du, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}

	return &httpRequestService{
		queue:   make(chan *Content),
		stop:    make(chan bool),
		limiter: tools.NewLimiter(du, 30),
	}
}
