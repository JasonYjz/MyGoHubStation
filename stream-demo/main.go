package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		mimeWriter := multipart.NewWriter(w)
		w.Header().Set("Content-Type", fmt.Sprintf("multipart/x-mixed-replace; boundary=%s", mimeWriter.Boundary()))
		conn, _, err := w.(http.Hijacker).Hijack()
		defer conn.Close()
		fmt.Fprintf(conn, "HTTP/1.1 200 OK")
		fmt.Fprintf(conn, "Content-Type: multipart/x-mixed-replace; boundary=myboundary")
		fmt.Fprintf(conn, "")
		c := make(chan []byte)
		go func() {
			for {
				frame, err := getFrame()
				if err != nil {
					log.Println(err)
					continue
				}
				c <- frame
			}
		}()
		for {
			select {
			case frame := <-c:
				fmt.Fprintf(conn, "--myboundary")
				fmt.Fprintf(conn, "Content-Type: image/jpeg")
				fmt.Fprintf(conn, "Content-Length: %d", len(frame))
				fmt.Fprintf(conn, "")
				conn.Write(frame)
			}
		}
	})
	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("listen server err:%s", err.Error())
}

func getFrame() ([]byte, error) {
	return []byte("xxx"), nil
}
