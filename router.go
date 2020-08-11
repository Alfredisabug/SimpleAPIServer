package main

import (
	servehandler "SimpleAPIServer/ServerHandler"
	"net/http"
)

func ServerInit() *http.Server {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	dataHandler := servehandler.DataHandler{}
	http.Handle("/data", &dataHandler)
	return &server
}
