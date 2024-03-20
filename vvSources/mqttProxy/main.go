package main

import (
	_ "bytes"
	"fmt"
	"io"
	"net"
)

const (
	mqttPort              = 1887
	mqttBroker            = "localhost"
	mqttPacketTypeMask    = 0xF0
	mqttControlPacketType = 0x30
)

func main() {
	listener, err := net.Listen("tcp", ":1883")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	mqttConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", mqttBroker, mqttPort))
	if err != nil {
		fmt.Println("Error connecting to MQTT broker:", err)
		return
	}
	defer mqttConn.Close()

	fmt.Println("TCP proxy listening on port:", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn, mqttConn)
	}
}

func handleClient(clientConn net.Conn, mqttConn net.Conn) {
	defer clientConn.Close()

	firstByte := make([]byte, 1)

	n, err := clientConn.Read(firstByte)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error reading first byte:", err)
		}
		return
	}

	if n != 1 || (firstByte[0]&mqttPacketTypeMask) != mqttControlPacketType {
		fmt.Println("Dropping non-MQTT packet")
		return
	}

	_, err = io.Copy(mqttConn, clientConn)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Error copying data:", err)
		}
		return
	}
	fmt.Printf("Forwarded %d bytes to MQTT broker\n", n)
}
