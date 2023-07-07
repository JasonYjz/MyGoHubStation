package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

const (
	diskPath = "/"
	B        = 1
	KB       = 1024 * B
	MB       = 1024 * KB
	GB       = 1024 * MB
)

func main() {
	path := os.Args[1]
	fmt.Printf("input path: %s \n", path)
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println("")
		return
	}

	total := uint64(fs.Blocks) * uint64(fs.Bsize)
	free := uint64(fs.Bavail) * uint64(fs.Bsize)
	used := (uint64(fs.Blocks) - uint64(fs.Bfree)) * uint64(fs.Bsize)

	fmt.Printf("Total: %f \n", formatTwoDecimalPlaces(float64(total)/GB))
	fmt.Printf("Free: %f \n", formatTwoDecimalPlaces(float64(free)/GB))
	fmt.Printf("Used: %f \n", formatTwoDecimalPlaces(float64(used)/GB))
	//d.Total = formatTwoDecimalPlaces(float64(total) / GB)
	//d.Free = formatTwoDecimalPlaces(float64(free) / GB)
	//d.Used = formatTwoDecimalPlaces(float64(used) / GB)
	//d.Unit = "GB"

	if (used + free) == 0 {
		//d.UsedRate = 0
		fmt.Printf("UsedRate: %f \n", 0.0)
	} else {
		fmt.Printf("UsedRate: %f \n", formatTwoDecimalPlaces((float64(used)/float64(used+free))*100.0))
	}
}

func formatTwoDecimalPlaces(value float64) float64 {
	ret, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	if err != nil {
		fmt.Println("failed to format two decimal places, use original value.")
		return value
	}
	return ret
}
