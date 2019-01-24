package method

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/charlesfan/go-tcp-server/service"
	"github.com/charlesfan/go-tcp-server/tcpmanager"
)

func PhotoHandler(c *tcpmanager.Context) {
	defer c.Clear()
	m := c.GetMessage()
	s := service.PixabayService
	p, err := parameters(m)
	if err != nil {
		c.Conn.Write([]byte(err.Error()))
	} else {
		s.GetPhotos(c.Conn, p)
	}
}

func VideoHandler(c *tcpmanager.Context) {
	defer c.Clear()
	m := c.GetMessage()
	s := service.PixabayService
	p, err := parameters(m)
	if err != nil {
		c.Conn.Write([]byte(err.Error()))
	} else {
		s.GetVideos(c.Conn, p)
	}
}

func PixabayAliveHandler(c *tcpmanager.Context) {
	defer c.Clear()
	s := service.PixabayService
	ok := s.Alive()
	c.Conn.Write([]byte(strconv.FormatBool(ok) + "\n"))
}

func parameters(s string) (string, error) {
	if s == "all" {
		return "", nil
	}
	p := url.Values{}
	a := strings.Split(s, " ")
	for _, v := range a {
		aa := strings.Split(v, "=")
		if string(v[0]) != "-" || len(aa) < 2 {
			return "", errors.New("Error Input\n")
		}
		k := strings.Replace(aa[0], "-", "", -1)
		p.Add(k, aa[1])
	}
	return p.Encode(), nil
}
