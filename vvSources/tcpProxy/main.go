package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func main() {
	proxyPort := ":1883"

	listener, err := net.Listen("tcp", proxyPort)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		downstreamConn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func(downstreamConn net.Conn) {
			defer downstreamConn.Close()

			isHTTPRequest, conn, err := peekHTTRequest(downstreamConn)
			if err != nil {
				return
			}

			if isHTTPRequest {
				fmt.Println("[+] Detected HTTP request, dropping packet")
				return
			}

			upstreamConn, err := net.Dial("tcp", ":1887")
			if err != nil {
				fmt.Println("Error connecting to upstream server:", err)
				return
			}
			defer upstreamConn.Close()

			go func() {
				_, err := io.Copy(upstreamConn, conn)
				if err != nil {
					fmt.Println("Error copying from downstream to upstream:", err)
				}
			}()
			_, err = io.Copy(downstreamConn, upstreamConn)
			if err != nil {
				fmt.Println("Error copying from upstream to downstream:", err)
			}
		}(downstreamConn)
	}
}

type bufferedConn struct {
	r    *bufio.Reader
	conn net.Conn
}

func newBufferedConn(c net.Conn) bufferedConn {
	return bufferedConn{bufio.NewReader(c), c}
}

func newBufferedConnSize(c net.Conn, n int) bufferedConn {
	return bufferedConn{bufio.NewReaderSize(c, n), c}
}

func (b bufferedConn) Peek(n int) ([]byte, error) {
	return b.r.Peek(n)
}

func (b bufferedConn) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b bufferedConn) Close() error {
	return b.conn.Close()
}

//func isHTTPRequest(conn net.Conn) bool {
//	const peekSize = 1024
//	bufCon := newBufferedConn(conn)
//	peek, _ := bufCon.Peek(32)
//	defer bufCon.Close()
//
//	if strings.Contains(string(peek), "HTTP") {
//		return true
//	}
//
//	return false
//}

type readOnlyConn struct {
	reader io.Reader
}

func (conn readOnlyConn) Read(p []byte) (int, error)         { return conn.reader.Read(p) }
func (conn readOnlyConn) Write(p []byte) (int, error)        { return 0, io.ErrClosedPipe }
func (conn readOnlyConn) Close() error                       { return nil }
func (conn readOnlyConn) LocalAddr() net.Addr                { return nil }
func (conn readOnlyConn) RemoteAddr() net.Addr               { return nil }
func (conn readOnlyConn) SetDeadline(t time.Time) error      { return nil }
func (conn readOnlyConn) SetReadDeadline(t time.Time) error  { return nil }
func (conn readOnlyConn) SetWriteDeadline(t time.Time) error { return nil }

func peekHTTRequest(conn net.Conn) (bool, io.Reader, error) {
	peekedBytes := new(bytes.Buffer)
	ok, err := isHTTPRequest(io.TeeReader(conn, peekedBytes))
	if err != nil {
		return false, nil, err
	}
	return ok, io.MultiReader(peekedBytes, conn), nil
}

func isHTTPRequest(reader io.Reader) (bool, error) {
	buf := make([]byte, 4096)
	readOnly := &readOnlyConn{reader: reader}
	_, err := readOnly.Read(buf)
	if err != nil {
		return false, err
	}

	if strings.Contains(string(buf), "HTTP") {
		return true, nil
	}

	return false, nil
}
