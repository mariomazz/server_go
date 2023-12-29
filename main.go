package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {
	port := os.Getenv("HTTP_PORT")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/ipaddress", getIpAddress)
	fmt.Printf("Server Started port: " + port + "\n")

	err := http.ListenAndServe(":" + port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

