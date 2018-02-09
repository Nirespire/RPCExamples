package main

import (
	"internal/haberdasherserver"

	"net/http"
	"rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
