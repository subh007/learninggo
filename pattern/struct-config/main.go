package main

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func getDefaults() *Opts {
	return &Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

type OptsFn func(*Opts)

func NewServer(fns ...OptsFn) *Server {
	defaults := getDefaults()
	for _, fn := range fns {
		fn(defaults)
	}
	return &Server{
		Opts: *defaults,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withMaxConn(maxconn int) OptsFn {
	return func(o *Opts) {
		o.maxConn = maxconn
	}
}

// This is example for the struct configuration pattern
func main() {
	server := NewServer(withMaxConn(100), withTls)
	fmt.Printf("Server %+v\n", server)
	fmt.Printf("maxcon %d", server.maxConn)
}
