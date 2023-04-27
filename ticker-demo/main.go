package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * time.Duration(3))
	go func() {
		for range ticker.C {
			fmt.Println("ticker is triggered.")
		}
	}()

	time.Sleep(1 * time.Minute)
}
