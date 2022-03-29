package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("Connection time Out")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "\nI heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code go here!")
}

func main() {
	ll, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ll.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}
