package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
	// "math/rand"

	"github.com/joho/godotenv"
)

const WIDTH = 300
const HEIGHT = 200

func runFfmpeg() *exec.Cmd {
	cmd := exec.Command(
		"ffmpeg",
		"-f", "rawvideo",
		"-video_size", fmt.Sprintf("%dx%d", WIDTH, HEIGHT),
		"-pixel_format", "rgb8",
		"-use_wallclock_as_timestamps", "1",
		// "-i", "unix:/tmp/go.sock",
		"-i", "-",
		"-c:v", "libx264",
		"-r", "60",
		"-f", "flv",
		os.Getenv("TARGET_OUTPUT"),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func spam_frames(ffmpeg io.Writer) {
	w, h := WIDTH, HEIGHT
	buf := make([]byte, w*h)

	println("Spamming now")

	shift := 0
	for {
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {

				if x < shift {
					buf[x+y*w] = byte((y / 0xFF) * (h / 0xFF))
				} else {
					buf[x+y*w] = 0xAF
				}

			}
		}

		shift += 1
		if shift > w {
			shift = 0
		}
		ffmpeg.Write(buf)

		time.Sleep(16 * time.Millisecond)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd := runFfmpeg()
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	println(stdin)

	go func () {
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	spam_frames(stdin)
}
