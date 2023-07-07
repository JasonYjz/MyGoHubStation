package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	fmt.Println("listen at :58080")
	http.HandleFunc("/ws", websocketHandler)
	http.ListenAndServe(":58080", nil)
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocketupgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()
	// 启动 FFmpeg 进行转码
	cmd := exec.Command("ffmpeg",
		//"-i", "input.mp4", // 输入文件
		"-rtsp_transport", "tcp",
		"-i", "rtsp://admin:aidlux123@192.168.110.234:554/h264/ch1/main/av_stream", // 输入文件
		//"-c:v", "libx264",
		//"-preset", "ultrafast",
		//"-tune", "zerolatency",
		"-r", "20",
		"-c:a", "copy",
		"-vcodec", "h264",
		//"-max_delay", "100",
		"-f", "flv",
		"pipe:1")

	fmt.Println("start to covert stream by ffmpeg")
	// 通过管道连接 FFmpeg 的标准输出和 WebSocket 的写入端
	ffmpegOutput, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Failed to create FFmpeg output pipe:", err)
		return
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("Failed to start FFmpeg:", err)
		return
	}
	defer cmd.Wait()
	go func() {
		// 从 FFmpeg 的标准输出读取数据并发送到 WebSocket 连接
		fmt.Println("loop to read the converted stream.")

		buf := make([]byte, 8192)
		num := 0
		for {
			n, err := ffmpegOutput.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Println("Error reading FFmpeg output:", err)
				}
				break
			}

			fmt.Printf("%s:read [%d] length data.\n", time.Now().Format("2006-01-02T15:04:05.000"), n)
			select {
			case <-time.Tick(1 * time.Second):
				fmt.Printf("====================read %d frames in 1s", num)
				num = 0
			default:
				num++
			}
			err = conn.WriteMessage(websocket.BinaryMessage, buf[:n])
			if err != nil {
				log.Println("Failed to write WebSocket message:", err)
				break
			}
		}
	}()
	// 监听 WebSocket 连接关闭事件，关闭 FFmpeg 进程_, _, err = conn.ReadMessage()
	if err != nil {
		fmt.Println("WebSocket connection closed:", err)
		cmd.Process.Kill()
	}

	fmt.Println("the process is end")
}
