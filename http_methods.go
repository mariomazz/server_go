package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"net"
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

func GetOutboundIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP.String();
}