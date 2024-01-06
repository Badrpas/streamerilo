package main

import (
	"image"
	"io"
	"math"
	"time"

	"github.com/fogleman/gg"
)

var frame_times [WIDTH]int64

func spam_frames(ffmpeg io.Writer) {
	frame_step := time.Duration(math.Trunc(1000000/SIMULATION_FPS)) * time.Microsecond

	println("Spamming now")
	im := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	ctx := gg.NewContextForRGBA(im)

	for {
		start := time.Now()

		for y := 0; y < HEIGHT; y++ {
			for x := 0; x < WIDTH; x++ {

			}
		}

        draw_frame(ctx)

		ffmpeg.Write(im.Pix)
		runtime := time.Since(start)
		frame_times[shift] = runtime.Milliseconds()
		var delta time.Duration = frame_step - runtime
		if delta < 0 {
			delta = 0
		}

		time.Sleep(delta)
	}

}

