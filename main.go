package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {

	var listener net.Listener
	var err error

	port := ":80"

	listener, err = net.Listen("tcp", port)
	if err != nil {
		port = ":8000"
		listener, err = net.Listen("tcp", port)
		if err != nil {
			log.Panic("Failed to listen on port 80 or 8000")
		}
	}
	log.Printf("Listening on port %s", port)

	http.HandleFunc("/", Hello)
	log.Fatal(http.Serve(listener, nil))
}
