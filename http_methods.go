package main

import (
	"fmt"
	"net/http"
	"io"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my server!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getIpAddress(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / getIpAddress\n")
	io.WriteString(w, "Id Address => "+r.Host+"\n")
}