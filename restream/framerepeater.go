package main

import (
	"math"
	"time"
)


func repeat_buffer() {
	frame_step := time.Duration(math.Trunc(1000000/float64(fps))) * time.Microsecond
	shift := 0
	max_shift := 100
	dir := 1
	for {
	    start := time.Now()

	    if ffmpeg_stdin != nil {
	        ffmpeg_stdin.Write(buffer)
	    }

        if clients_connected == 0 {
            draw_rect(uint32(width / 2 + uint32(shift - max_shift / 2)), height / 2, 20, 20)
            shift += dir
            if shift == max_shift || shift == 0 {
                dir = -dir
            }
        }

		runtime := time.Since(start)

		var delta time.Duration = frame_step - runtime
		if delta < 0 {
			delta = 0
		}

		time.Sleep(delta)
	}
}
