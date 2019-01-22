package method

import (
	"fmt"

	"github.com/charlesfan/tcp-server/handler"
)

func QuitHandler(c *handler.Context) {
	fmt.Println("Quitting.")
	c.Conn.Write([]byte("Disconnection!!!\n"))
	c.Conn.Close()
}
