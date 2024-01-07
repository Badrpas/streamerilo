package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    listener, err := net.Listen(
        "tcp4",
        "127.0.0.1:3322",
    )
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()
    go repeat_buffer()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go handle_req(conn)
    }
}



