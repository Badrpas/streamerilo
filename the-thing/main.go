package main

import (
	"log"

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

	stream_to_remote := true
	cmd := GetFfmpegCmd(stream_to_remote, TARGET_FPS, WIDTH, HEIGHT)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	if !stream_to_remote {
		go RunFfplay()
	}

	spam_frames(stdin)
}
