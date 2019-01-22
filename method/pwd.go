package method

import (
	"github.com/charlesfan/tcp-server/handler"
)

func PwdHandler(c *handler.Context) {
	c.Conn.Write([]byte("Now, your state: " + c.MethodName + "\n"))
}
