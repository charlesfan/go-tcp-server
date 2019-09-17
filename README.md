# go-tcp-server
Practice tcp server with Golang.This tcp server takes in any request text per line to fetch images and videos by sending queries to an external API: [Pixabay](https://pixabay.com/api/docs/)
- [x] Handle text per line.
- [x] Simple tcp framwork to add new handler function in easy way.
- [x] Query to an external API.
- [x] Limit for the external API request: 30 requests per second.
- [x] Check the status of the external API server.
- [x] Client will disconnect when send 'quit' or timed out(120 second).
- [ ] Test all service and method.
- [ ] HTTP endpoint to display some statistics of this server. 

## Run server
1. Clone this code:
```sh
$ git clone https://github.com/charlesfan/go-tcp-server.git
```
2. Set API key:
```sh
$ export PKey=<your key>
```
3. Build server:
```sh
$ go build server.go
```
4. Run tcp server.go wiht port flag (default port: 3333):
```sh
$ ./server -port <port number>
```
## Test
```sh
$ make test
```
## How to use
Use telnet command or tcp clinet application to send message.
Provide two Method: photo and video
### Get photos
 
```
> photo -color=red -size=123
```
### Get videos
```
> photo -videos=large
```
We can get other serch parameters at [Pixabay's API document](https://pixabay.com/api/docs/)

### New method in server side
We can easy to add new method and handler function to tcp server:
1. Modify method/method.go
```golang
package method

import (
	"github.com/charlesfan/go-tcp-server/tcpmanager"
)

func Init() *tcpmanager.Engine {
	e := tcpmanager.New()

	e.NewMethod("photo", PhotoHandler, false)
        e.NewMethod("video", VideoHandler, false)
        e.NewMethod("pixabay", PixabayAliveHandler, true)
+       e.NewMethod("nameOfNewMethod", NewMethodHandler, true)

	return e
}
```
2. Create new file at method/:

#### method/newMethodHandler.go
```golang
package method

import (
  "github.com/charlesfan/go-tcp-server/tcpmanager"
)

func NewMethodHandler(c *tcpmanager.Context) {
  c.Conn.Write([]byte("Hello world!!!\n"))
}

```
3. Restart server
