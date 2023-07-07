package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	var s = http.Server{
		Addr: ":5858",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("receive a r from:", r.RemoteAddr, r.Header)
			time.Sleep(10 * time.Second)
			w.Write([]byte("ok"))
		}),
	}

	s.ListenAndServe()
}
