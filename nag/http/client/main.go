package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	server := flag.String("server", "localhost:8080", "HTTP Server URL")
	check := flag.Bool("check", false, "healthcheck flag")
	lookup := flag.String("lookup", "domain", "lookup data [mac, ip, domain]")
	flag.Parse()

	if flag.NFlag() != 1 && !(*check) {
		log.Println("must provide exactly one query argument")
		return
	}

	path := "/lookup"
	if *check {
		path = "/check"
	}

	addr, err := url.Parse("http://" + *server + path)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add(*lookup, flag.Arg(0))
	addr.RawQuery = params.Encode()
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Get(addr.String())
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
