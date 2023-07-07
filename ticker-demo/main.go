package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//ticker := time.NewTicker(time.Second * time.Duration(3))
	//go func() {
	//	for range ticker.C {
	//		fmt.Println("ticker is triggered.")
	//	}
	//}()
	//
	//time.Sleep(1 * time.Minute)

	strs := strings.Split("XXX_1.0.0_20230511_151213", "_")
	dir := filepath.Join(strs[0], strs[1], strs[2], strs[3][0:2], strs[3][2:4], strs[3][4:])
	fmt.Printf("dir: %s\n", dir)
}
