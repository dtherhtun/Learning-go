package main

import "fmt"

type OptFunc func(opts *Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

type Server struct {
	Opts
}

func NewServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := NewServer(withTLS, withMaxConn(5), withID("foo bar"))
	fmt.Printf("%+v\n", s)
}
