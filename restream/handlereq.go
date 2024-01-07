package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

var width uint32
var height uint32
var pixel_size uint32
var fps uint32

var buffer []byte

var ffmpeg_stdin io.Writer

var offset int = 0

var clients_connected int = 0

func handle_req(conn net.Conn) {
	clients_connected++
	header := make([]byte, 4*4)
	size, err := conn.Read(header)
	println("size is ", size)
	if err != nil {
		log.Fatal(err)
	}
	width, height, pixel_size = binary.LittleEndian.Uint32(header), binary.LittleEndian.Uint32(header[4:]), binary.LittleEndian.Uint32(header[8:])
	fps = binary.LittleEndian.Uint32(header[12:])

	expected_size := int(width * height * pixel_size)
	println("w ", width, " h ", height, " p ", pixel_size)
	if len(buffer) != int(expected_size) {
		println("Setting buffer size to ", expected_size)
		buffer = make([]byte, expected_size)
		if ffmpeg_stdin == nil {
			cmd := GetFfmpegCmd(true, float64(fps), int(width), int(height))
			ffmpeg_stdin, err = cmd.StdinPipe()
			if err != nil {
				log.Fatal(err)
			}
			go func() {
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
			}()
		}
	}

	local := make([]byte, expected_size)
	offset = 0

	for {
		received, err := conn.Read(local)
		// println("Read ", received, " bytes")
		if err != nil {
			println(err)
			println("Did client disconnect?")
			clients_connected--
			break
		}

		idx := 0
		for idx < received {
			expected := expected_size - offset
			copy_size := expected
			if expected > received-idx {
				copy_size = received - idx
			}
			// println("idx ", idx, " received ", received, " expected ", expected)
			for i := 0; i < copy_size; i++ {
				buffer[offset+i] = local[idx+i]
			}
			idx += copy_size
			offset += copy_size
			offset %= expected_size
		}

	}

}

func draw_rect(cx, cy, w, h uint32, r, g, b byte) {
	const S = 20
	for y := cy - h/2; y < cy+h/2; y++ {
		for x := cx - w/2; x < cx+w/2; x++ {
			buffer[x*pixel_size + y*width*pixel_size + 0] = r
			buffer[x*pixel_size + y*width*pixel_size + 1] = g
			buffer[x*pixel_size + y*width*pixel_size + 2] = b
			buffer[x*pixel_size + y*width*pixel_size + 3] = 0xFF
		}
	}

}
