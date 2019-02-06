package service_test

import (
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/charlesfan/go-tcp-server/service"
	"github.com/stretchr/testify/assert"
)

type Addr struct {
	NetworkString string
	AddrString    string
}

func (a Addr) String() string {
	return a.AddrString
}

func (a Addr) Network() string {
	return a.NetworkString
}

type FakeConn struct{}

func (c FakeConn) Close() error { return nil }
func (e FakeConn) Read(data []byte) (n int, err error) {
	n = 128
	err = nil
	return
}

func (e FakeConn) Write(data []byte) (n int, err error) {
	n = 128
	err = nil
	return
}

func (e FakeConn) LocalAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e FakeConn) RemoteAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e FakeConn) SetDeadline(t time.Time) error      { return nil }
func (e FakeConn) SetReadDeadline(t time.Time) error  { return nil }
func (e FakeConn) SetWriteDeadline(t time.Time) error { return nil }

var reqCache []time.Time

type HttpRequestTestCaseSuite struct {
	service    service.HttpRequestServicer
	testServer *httptest.Server
	count      uint32
	first      time.Time
}

func setupHttpRequestTestCaseSuite(t *testing.T) (HttpRequestTestCaseSuite, func(t *testing.T)) {
	s := HttpRequestTestCaseSuite{
		service: service.NewHttpRequestService(),
	}

	SimpleHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		atomic.AddUint32(&s.count, 1)
		switch atomic.LoadUint32(&s.count) {
		case 1:
			s.first = time.Now()
		case 30:
			diff := time.Now().Sub(s.first)
			assert.False(t, diff.Seconds() > float64(0.999))
		case 31:
			diff := time.Now().Sub(s.first)
			assert.True(t, diff.Seconds() > float64(0.99))
		}
	}

	s.testServer = httptest.NewServer(http.HandlerFunc(SimpleHandler))

	return s, func(t *testing.T) {
		s.service.Stop()
	}
}

func TestHttpRequestService(t *testing.T) {
	s, teardownTestCase := setupHttpRequestTestCaseSuite(t)
	defer teardownTestCase(t)

	go s.service.Run()
	for i := 0; i < 40; i++ {
		c := &service.Content{
			Conn:   &FakeConn{},
			Method: "GET",
			Url:    s.testServer.URL,
		}

		s.service.Put(c)
	}
}
