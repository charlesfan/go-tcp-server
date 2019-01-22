package method

import (
	"net/url"
	"strings"

	"github.com/charlesfan/go-tcp-server/tcpmanager"
)

func PhotoHandler(c *tcpmanager.Context) {
	defer c.Clear()
	m := c.GetMessage()
	if string(m[0]) != "-" {
		c.Conn.Write([]byte("Error Input\n"))
	} else {
		p := parameters(m)
		c.Conn.Write([]byte(p + "\n"))
	}
}

func parameters(s string) string {
	p := url.Values{}
	a := strings.Split(s, " ")
	for i, v := range a {
		if string(v[0]) == "-" {
			v := strings.Replace(v, "-", "", -1)
			p.Add(v, a[i+1])
		}
	}
	return p.Encode()
}
