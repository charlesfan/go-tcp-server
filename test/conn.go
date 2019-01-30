package test

import (
	"net"
	"time"
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
