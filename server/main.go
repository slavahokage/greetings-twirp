package main

import (
	"helloserver"
	"helloservice"
	"net/http"
)

func main() {
	server := &helloserver.Server{} // implements Haberdasher interface
	twirpHandler := hello.NewGreetingServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}