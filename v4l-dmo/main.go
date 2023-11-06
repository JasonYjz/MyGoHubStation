package main

import (
	"fmt"
	"github.com/vladimirvivien/go4vl/device"
	"github.com/vladimirvivien/go4vl/v4l2"
	"io"
	"log"
	"net/http"
)

var frames <-chan []byte

func main() {
	devName := "/dev/video2"
	// create device
	dev, err := device.Open(
		devName,
		device.WithIOType(v4l2.IOTypeMMAP),
		device.WithPixFormat(v4l2.PixFormat{
			PixelFormat: getFormatType(format),
			Width:       uint32(width),
			Height:      uint32(height),
		}), device.WithFPS(uint32(frameRate)),
	)

}

func serveVideoStream(w http.ResponseWriter, req *http.Request) {
	// Start HTTP Response
	const boundaryName = "Yt08gcU534c0p4Jqj0p0"
	// send multi-part header
	w.Header().Set(
		"Content-Type",
		fmt.Sprintf(
			"multipart/x-mixed-replace; boundary=%s", boundaryName,
		),
	)
	w.WriteHeader(http.StatusOK)
	for frame := range frames {
		// start boundary
		io.WriteString(w, fmt.Sprintf("--%s\n", boundaryName))
		io.WriteString(w, "Content-Type: image/jpeg\n")
		io.WriteString(w, fmt.Sprintf(
			"Content-Length: %d\n\n", len(frame)))
		if _, err := w.Write(frame); err != nil {
			log.Printf("failed to write mjpeg image: %s", err)
			return
		}

		// close boundary
		if _, err := io.WriteString(w, "\n"); err != nil {
			log.Printf("failed to write boundary: %s", err)
			return
		}
	}
}
