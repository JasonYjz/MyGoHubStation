package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"log"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Print(">>> 打开摄像头失败！\n")
		return
	}
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("没有读取到图像！")
			return
		}
		if img.Empty() {
			fmt.Printf("空图像！")
			return
		}

		window.IMShow(img)
		key := window.WaitKey(1)
		if key == 27 {
			return
		}
	}
}

func Test() {
	capture, err := gocv.VideoCaptureFile("../../data/drop.avi")
	if err != nil {
		log.Panic("Can not find video")
		return
	}
	defer capture.Close()

	window := gocv.NewWindow("Video")
	defer window.Close()

	mat := gocv.NewMat()

	for {
		ok := capture.Read(&mat)
		if !ok {
			log.Println("Reached the end of the video")
			break
		}

		window.IMShow(mat)

		key := window.WaitKey(1)
		if key == 27 {
			break
		}
	}
}
