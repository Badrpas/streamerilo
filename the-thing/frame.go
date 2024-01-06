package main

import (
	"fmt"
	"time"

	"github.com/fogleman/gg"
)

var shift int = 0

func draw_frame(ctx *gg.Context) {
	if shift == 0 {
	    ctx.Clear()
	}
	ctx.SetRGB(0, 0.1, 0.3)
	ctx.DrawRectangle(0, 0, WIDTH, 240)
	ctx.Fill()
	ctx.SetRGB(1, 0.1, 0.3)
	// ctx.DrawCircle(-100. + (float64(shift) / float64(WIDTH)) * float64(WIDTH + 200.), 100, 100)
	// ctx.Fill()

	current := time.Now()
	ctx.DrawString(current.Format("2006-01-02 15:04:05.000000"), 10, 220)
	ctx.DrawString(fmt.Sprintf("%.2f", avg_last_frame_duration(60)), 10, 230)

	// for i := 0; i < WIDTH; i++ {
	i := (shift + WIDTH ) % WIDTH
	time := frame_times[i]
	ctx.MoveTo(float64(i), HEIGHT)
	ctx.LineTo(float64(i), float64(HEIGHT-time))
	// }
	ctx.Stroke()

	shift += 1
	if shift == WIDTH {
		shift = 0
	}

}


func avg_last_frame_duration(n int) float64 {
    sum := 0.
    for i:=0; i < n; i++ {
        sum += float64(frame_times[(shift + WIDTH - i) % WIDTH])
    }
    return sum / float64(n)
}
