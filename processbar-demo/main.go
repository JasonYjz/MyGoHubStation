package main

import (
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	TestSpinnerCustom()
	ExampleOptionThrottle()

	time.Sleep(2 * time.Second)
}

func TestSpinnerCustom() {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetWidth(10),
		progressbar.OptionSetDescription("indeterminate spinner"),
		progressbar.OptionShowIts(),
		progressbar.OptionShowCount(),
		progressbar.OptionSpinnerCustom([]string{"↖", "↗", "↘", "↙"}),
	)
	bar.Reset()
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		err := bar.Add(1)
		if err != nil {
		}
	}
}

func ExampleOptionThrottle() {
	bar := progressbar.NewOptions(100,
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(100*time.Millisecond))
	bar.Reset()
	bar.Add(5)
	time.Sleep(150 * time.Millisecond)
	bar.Add(5)
	bar.Add(10)
	// Output:
	// 10% |█         |  [0s:1s]
}
