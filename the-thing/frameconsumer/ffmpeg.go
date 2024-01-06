package frameconsumer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetFfmpegCmd(use_env bool, target_fps float64, width, height int) *exec.Cmd {
	target_fps_str := fmt.Sprintf("%.0f", target_fps)

	cmd := exec.Command(
		"ffmpeg",
		"-f", "rawvideo",
		"-video_size", fmt.Sprintf("%dx%d", width, height),
		"-pixel_format", "rgba",
		"-use_wallclock_as_timestamps", "1",
		"-i", "-",
		"-c:v", "libx264",
		"-r", target_fps_str,
		"-f", "flv",
		os.Getenv("TARGET_OUTPUT"),
	)

	if !use_env {
		cmd = exec.Command(
			"ffmpeg",
			"-f", "rawvideo",
			"-video_size", fmt.Sprintf("%dx%d", width, height),
			"-pixel_format", "rgba",
			"-use_wallclock_as_timestamps", "1",
			"-i", "-",
			"-c:v", "libx264",
			"-r", target_fps_str,
			"-f", "flv",
			"-listen", "1",
			"unix:/tmp/out.sock",
		)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}


func RunFfplay() {
	cmd := exec.Command(
		"ffplay",
		"unix:/tmp/out.sock",
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
