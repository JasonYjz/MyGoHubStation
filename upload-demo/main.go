package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "upload-demo/static/index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// 接收文件
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("File retrieval error:", err)
		return
	}
	defer file.Close()

	// 创建目标文件
	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Println("File creation error:", err)
		return
	}
	defer out.Close()

	// 获取文件大小
	fileStat, err := out.Stat()
	if err != nil {
		fmt.Println("File stat error:", err)
		return
	}
	fileSize := fileStat.Size()

	// 创建缓冲区
	buffer := make([]byte, 1024)

	// 读取并写入文件，同时发送上传进度到前端
	var uploadedBytes int64
	for {
		// 从文件读取数据
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("File read error:", err)
			}
			break
		}

		// 写入文件
		_, err = out.Write(buffer[:n])
		if err != nil {
			fmt.Println("File write error:", err)
			break
		}

		// 更新已上传的字节数
		uploadedBytes += int64(n)

		// 计算上传进度
		progress := float64(uploadedBytes) / float64(fileSize) * 100

		// 发送进度信息给前端
		err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%.2f%%", progress)))
		if err != nil {
			fmt.Println("WebSocket write error:", err)
			break
		}
	}

	// 上传完成，发送完成消息给前端
	err = conn.WriteMessage(websocket.TextMessage, []byte("Upload completed"))
	if err != nil {
		fmt.Println("WebSocket write error:", err)
	}
}
