package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	inputURL := "rtsp://input_stream_url"
	outputURL := "rtmp://output_stream_url"

	cmd := exec.Command("ffmpeg",
		"-i", inputURL,
		"-c:v", "libx264",
		"-preset", "ultrafast",
		"-tune", "zerolatency",
		"-c:a", "aac",
		"-f", "flv",
		outputURL,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transcoding completed.")
}
