package main

import (
	"fmt"
	"io"
	"net"
	"strings"
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

			buf := make([]byte, 4096)
			n, err := downstreamConn.Read(buf)
			if err != nil {
				fmt.Println("Error reading from downstream connection:", err)
				return
			}

			upstreamConn, err := net.Dial("tcp", ":1887")
			if err != nil {
				fmt.Println("Error connecting to upstream server:", err)
				return
			}
			defer upstreamConn.Close()
			if strings.Contains(string(buf), "HTTP") {
				fmt.Println("[+] Detected HTTP request, dropping packet")
				return
			}

			_, err = upstreamConn.Write(buf[:n])
			if err != nil {
				fmt.Println("Error writing data to upstream connection:", err)
				return
			}

			go func() {
				_, err := io.Copy(upstreamConn, downstreamConn)
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
