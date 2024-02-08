package frameconsumer

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)


func ConnectRestreamer(width, height, pixel_size, fps uint32) io.Writer {
    addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:3322")
    if err != nil {
        log.Fatal(err)
    }
    conn, err := net.DialTCP("tcp4", nil, addr)
    if nerr, ok := err.(*net.OpError); ok && nerr.Op == "dial" && strings.Contains(nerr.Error(), "connection refused") {
        fmt.Printf("Restreamer Connection refused; Retrying in 1 second\n");
        time.Sleep(time.Second)
        conn, err = net.DialTCP("tcp4", nil, addr)
    }
    if err!= nil {
        log.Fatal(err)
    }

	header := make([]byte, 16)
    binary.LittleEndian.PutUint32(header, width)
    binary.LittleEndian.PutUint32(header[4:], height)
    binary.LittleEndian.PutUint32(header[8:], pixel_size)
    binary.LittleEndian.PutUint32(header[12:], fps)
    conn.Write(header)
    return conn
}

