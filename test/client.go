package test

import (
	"bufio"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SendMessage(t *testing.T, host string, message string) string {
	conn, err := net.Dial("tcp", host)
	defer conn.Close()
	assert.NoError(t, err)
	defer conn.Close()
	fmt.Fprintf(conn, message+"\n")
	text, _ := bufio.NewReader(conn).ReadString('\n')
	return text
}
