package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v\n", "Well, I hope!")
		conn.Close()
	}
}
