package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
    "fmt"
)

func echoServer(c net.Conn) {
	for {
		w, h := (200), (100)
		// c.Write([]byte{w, h})
		buf := make([]byte, w*h)

		print(len(buf))

		shift := byte(0)
		println("Spamming now")
		for {
			counter := byte(shift)
			for i:=0; i < int(w)*int(h); i++ {
				buf[i] = counter
				counter += 1
			}

			shift += 1
			c.Write(buf)
		}

		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got size:", len(data))
	}
}

func main() {
    fmt.Println("Hello, world.")

	ln, err := net.Listen("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(ln, sigc)


	for {
		println("Waiting for connection")
		fd, err := ln.Accept()
		println("Someone just got accepted!")
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		go echoServer(fd)
	}

}


