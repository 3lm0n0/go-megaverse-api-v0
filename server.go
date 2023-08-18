package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultPort = "3000"
	idleTimeout       = 30 * time.Second
	writeTimeout      = 180 * time.Second
	readHeaderTimeout = 10 * time.Second
	readTimeout       = 10 * time.Second
)

type Server interface {
	Start() error
}

type ServerParams struct {
	listenAddress string
}

func NewServer(listenAddress string) Server {
	return &ServerParams{
		listenAddress: listenAddress,
	}
}

func(sp *ServerParams) Start() error {
	if sp.listenAddress == "" {
		sp.listenAddress = defaultPort
	}
	server := &http.Server{
		Addr:    "0.0.0.0:" + sp.listenAddress,

		IdleTimeout:       idleTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
	}
	fmt.Println("server runnig at: ", server.Addr)
	return server.ListenAndServe()
}