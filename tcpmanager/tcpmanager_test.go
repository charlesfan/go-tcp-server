package tcpmanager

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/charlesfan/go-tcp-server/test"
)

type TcpmanagerTestCaseSuite struct {
	engine *Engine
	host   string
}

func setupTcpmanagerTestCaseSuite(t *testing.T) (TcpmanagerTestCaseSuite, func(t *testing.T)) {
	s := TcpmanagerTestCaseSuite{
		host: "localhost:3333",
	}
	e, err := NewListener(s.host)
	assert.NoError(t, err)
	s.engine = e
	//Handler function
	var f HandlerFunc = func(c *Context) {
		c.Conn.Write([]byte(c.MethodName + "\n"))
	}

	var f2 HandlerFunc = func(c *Context) {
		c.Conn.Write([]byte(c.message + "\n"))
	}

	var quitf HandlerFunc = func(c *Context) {
		c.Conn.Close()
	}
	//Add new method
	s.engine.NewMethod("pwd", f, true)
	s.engine.NewMethod("test", f2, false)
	s.engine.NewMethod("quit", quitf, true)

	return s, func(t *testing.T) {
		s.engine.Listener.Close()
	}
}

func TestAddMethod(t *testing.T) {
	s, teardownTestCase := setupTcpmanagerTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name          string
		methodName    string
		err           bool
		root          bool
		handler       HandlerFunc
		handlerLen    int
		setupTestCase test.SetupSubTest
	}{
		{
			name:       "success",
			methodName: "testHandler",
			err:        false,
			root:       false,
			handlerLen: 2,
			handler: func(c *Context) {
				c.Conn.Write([]byte(c.message + "\n"))
			},
			setupTestCase: test.EmptySubTest(),
		},
		{
			name:       "success with no paramater function",
			methodName: "testHandler2",
			err:        false,
			root:       true,
			handlerLen: 1,
			handler: func(c *Context) {
				c.Conn.Write([]byte(c.message + "\n"))
			},
			setupTestCase: test.EmptySubTest(),
		},
		{
			name:       "duplicate",
			methodName: "test",
			err:        true,
			root:       true,
			handler: func(c *Context) {
				c.Conn.Write([]byte(c.message + "\n"))
			},
			setupTestCase: test.EmptySubTest(),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupTestCase(t)
			defer teardownSubTest(t)

			if tc.err {
				assert.Panics(t, func() { s.engine.NewMethod(tc.methodName, tc.handler, tc.root) })
			} else {
				s.engine.NewMethod(tc.methodName, tc.handler, tc.root)
				assert.NotNil(t, s.engine.handlerMap[tc.methodName])
				assert.Equal(t, s.engine.handlerMap[tc.methodName].Len(), tc.handlerLen)
			}
		})
	}
}

func TestConn(t *testing.T) {
	s, teardownTestCase := setupTcpmanagerTestCaseSuite(t)
	defer teardownTestCase(t)

	tt := []struct {
		name          string
		methodName    string
		parameters    string
		wantText      string
		setupTestCase test.SetupSubTest
	}{
		{
			name:       "succes",
			methodName: "pwd",
			parameters: "",
			wantText:   "root\n",
			setupTestCase: func(t *testing.T) func(t *testing.T) {
				return func(t *testing.T) {
				}
			},
		},
		{
			name:       "Test method with parameters",
			methodName: "test",
			parameters: "parameters",
			wantText:   "parameters\n",
			setupTestCase: func(t *testing.T) func(t *testing.T) {
				return func(t *testing.T) {
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			teardownSubTest := tc.setupTestCase(t)
			defer teardownSubTest(t)

			go func() {
				m := strings.Join([]string{tc.methodName, tc.parameters}, " ")

				text := test.SendMessage(t, s.host, m)
				defer test.SendMessage(t, s.host, "quit")
				assert.Equal(t, text, tc.wantText)
			}()

			for {
				conn, err := s.engine.Listener.Accept()
				assert.NoError(t, err)

				s.engine.handleConnection(conn)
				return
			}
		})
	}
}
