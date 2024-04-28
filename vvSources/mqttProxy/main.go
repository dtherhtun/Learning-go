package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	proxyPort := ":1883"
	upstreamServer := ":1887"

	listener, err := net.Listen("tcp", proxyPort)
	if err != nil {
		log.Fatalf("Error starting proxy server: %v", err)
	}
	defer listener.Close()

	log.Printf("Proxy server started on %s\n", proxyPort)

	// Set up signal handler for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down...")
		listener.Close()
	}()

	for {
		downstreamConn, err := listener.Accept()
		if err != nil {
			// Check if the error is due to listener being closed during shutdown
			if strings.Contains(err.Error(), "use of closed network connection") {
				log.Println("Listener closed, exiting...")
				return
			}
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(downstreamConn, upstreamServer)
	}
}

func handleConnection(downstreamConn net.Conn, upstreamServer string) {
	defer downstreamConn.Close()

	reader := bufio.NewReader(downstreamConn)
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading from downstream connection: %v", err)
		return
	}

	if strings.Contains(firstLine, "HTTP") {
		log.Println("[+] Detected HTTP request, dropping packet")
		return
	}

	upstreamConn, err := net.Dial("tcp", upstreamServer)
	if err != nil {
		log.Printf("Error connecting to upstream server: %v", err)
		return
	}
	defer upstreamConn.Close()

	_, err = upstreamConn.Write([]byte(firstLine))
	if err != nil {
		log.Printf("Error writing data to upstream connection: %v", err)
		return
	}

	go io.Copy(upstreamConn, downstreamConn)
	_, err = io.Copy(downstreamConn, upstreamConn)
	if err != nil {
		log.Printf("Error copying from upstream to downstream: %v", err)
	}
}
