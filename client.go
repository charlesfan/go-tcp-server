package client

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 3333, "The port to connect to; defaults to 3333.")

func main() {
	flag.Parse()

	dest := *host + ":" + strconv.Itoa(*port)
	fmt.Printf("Connecting to %s...\n", dest)

	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}

	go readConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
}

func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		ok := scanner.Scan()
		text := scanner.Text()

		switch {
		case text == "Disconnection!!!":
			fmt.Println("\b\bQuit...")
			os.Exit(0)
		}
		fmt.Printf("\b** %s\n> ", text)

		if !ok {
			fmt.Println("Reached EOF on server connection.")
			break
		}
	}
}
