package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func handleClient(client net.Conn) {
	defer client.Close()

	buf := make([]byte, 4096)

	_, err := client.Read(buf)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	if strings.Contains(string(buf), "HTTP") {
		fmt.Println("[+] Detected HTTP request, dropping packet")
		return
	}

	targetHost := "localhost"
	targetPort := "1887"
	targetConn, err := net.Dial("tcp", targetHost+":"+targetPort)
	if err != nil {
		fmt.Println("Error connecting to target server:", err)
		return
	}
	defer targetConn.Close()
	go io.Copy(targetConn, client)
	io.Copy(client, targetConn)
}

func main() {

	proxyPort := ":1883"

	listener, err := net.Listen("tcp", proxyPort)
	if err != nil {
		fmt.Println("Error starting proxy server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("[+] Proxy server listening on port", proxyPort)

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(client)
	}
}

func Proxy() {
	proxyPort := ":1883"

	listener, err := net.Listen("tcp", proxyPort)
	if err != nil {
		fmt.Println(err)
	}

	for {
		downstreamCoon, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		buf := make([]byte, 4096)

		_, err = downstreamCoon.Read(buf)
		switch err {
		case nil:
			fmt.Println("nil")
		case io.EOF:
			fmt.Println("IO  EOF", err)
		default:
			fmt.Println("ERROR", err)
		}
		fmt.Println(string(buf))

		if strings.Contains(string(buf), "HTTP") {
			fmt.Println("[+] Detected HTTP request, dropping packet")

		}
		go func(downConn net.Conn) {
			defer downConn.Close()
			upstreamConn, err := net.Dial("tcp", ":1887")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer upstreamConn.Close()
			go io.Copy(upstreamConn, downConn)
			io.Copy(downConn, upstreamConn)
		}(downstreamCoon)
	}
}

// MQIsdp
