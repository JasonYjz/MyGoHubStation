package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start to process for panic-demo")
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
	}

	panic(1)

	fmt.Println("block to running always")
	select {}
}
