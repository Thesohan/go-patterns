package main

import (
	"fmt"
)

/*
we are building a server package that allows the user to configure a few options, such as the server ID,
maximum connections, and whether to use TLS encryption. We can use the options pattern to achieve this easily.
https://raufhm.hashnode.dev/gos-options-pattern
*/

type Options struct {
	ID      string
	MaxConn int64
	Tls     bool
}

type OptFunc func(*Options)

func DefaultOpts() Options {
	return Options{
		ID:      "01",
		MaxConn: 10,
		Tls:     false,
	}
}

func WithTls(opts *Options) {
	opts.Tls = true
}

func WithId(id string) OptFunc {
	return func(options *Options) {
		options.ID = id
	}
}

func WithMaxConn(mc int64) OptFunc {
	return func(options *Options) {
		options.MaxConn = mc
	}
}

type Server struct {
	Options
}

func NewServer(Opts ...OptFunc) *Server {
	opts := DefaultOpts()

	for _, fn := range Opts {
		fn(&opts)
	}

	return &Server{
		Options: opts,
	}
}

func main() {
	s := NewServer(WithTls, WithId("02"), WithMaxConn(100))
	fmt.Println(s)
}
