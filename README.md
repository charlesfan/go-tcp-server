# go-tcp-server
Practice tcp server with Golang.This tcp server takes in any request text per line and send a query to an external API: [Pixabay](https://pixabay.com/api/docs/)
- [x] Handle text per line.
- [x] Simple tcp framwork to add new handler function in easy way.
- [x] Query to an external API.
- [x] Limit for the external API request: 30 requests per second.
- [x] Check the status of the external API server.
- [ ] Test all service and method.
- [ ] HTTP endpoint to display some statistics of this server. 

## Install
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
4. Run tcp server.go:
```sh
$ ./server
```
## Test
```sh
$ make test
```
