package main

import (
	"io"
	"log"
	"os"

	. "thingy/frameconsumer"

	"github.com/joho/godotenv"
)

const WIDTH = 640
const HEIGHT = 480
const PIXEL_SIZE = 4

const SIMULATION_FPS float64 = 60
const TARGET_FPS float64 = SIMULATION_FPS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	use_restreamer := os.Getenv("USE_RESTREAMER") == "true"
	stream_to_remote := true

	var framereceiver io.Writer

	if use_restreamer {
		framereceiver = ConnectRestreamer(WIDTH, HEIGHT, PIXEL_SIZE, uint32(SIMULATION_FPS))
	} else {
		cmd := GetFfmpegCmd(stream_to_remote, TARGET_FPS, WIDTH, HEIGHT)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		framereceiver = stdin
		go func() {
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}()
		if !stream_to_remote {
			go RunFfplay()
		}
	}

	spam_frames(framereceiver)
}
