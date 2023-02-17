package main

import (
	"log"
	"net"

	"github.com/jsimonetti/rtnetlink/rtnl"
)

func main() {
	conn, err := rtnl.Dial(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	links, err := conn.Links()
	if err != nil {
		log.Fatal(err)
	}

	var loopBack *net.Interface
	for _, l := range links {
		if l.Name == "lo0" {
			loopBack = l
			log.Printf("Name: %s, Flags:%s\n", l.Name, l.Flags)
		}
	}

	err = conn.LinkDown(loopBack)
	if err != nil {
		log.Fatal(err)
	}

	loopBack, err = conn.LinkByIndex(loopBack.Index)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name: %s, Flags:%s\n", loopBack.Name, loopBack.Flags)

	err = conn.LinkUp(loopBack)
	if err != nil {
		log.Fatal(err)
	}
	loopBack, err = conn.LinkByIndex(loopBack.Index)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Name: %s, Flags:%s\n", loopBack.Name, loopBack.Flags)
}
