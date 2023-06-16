package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct {
	w      io.Writer
	donech chan struct{}
}

var tunnels = map[int]chan Tunnel{}

func main() {
	go func() {
		http.HandleFunc("/", handleRequest)
		log.Fatal(http.ListenAndServe(":3000", nil))
	}()

	ssh.Handle(func(session ssh.Session) {
		id := rand.Intn(math.MaxInt)
		tunnels[id] = make(chan Tunnel)

		fmt.Println("Tunnel ID ->", id)

		tunnel := <-tunnels[id]
		fmt.Println("tunnel is ready")

		_, err := io.Copy(tunnel.w, session)
		if err != nil {
			log.Fatal(err)
		}
		close(tunnel.donech)

		session.Write([]byte("we are DONE!!"))
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idstr)

	tunnel, ok := tunnels[id]
	if !ok {
		w.Write([]byte("tunnel not found"))
	}

	donech := make(chan struct{})
	tunnel <- Tunnel{
		w:      w,
		donech: donech,
	}

	<-donech
}
