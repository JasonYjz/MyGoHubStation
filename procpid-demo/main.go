package main

import (
	"fmt"
	"go-hub-station/procpid-demo/util"
)

func main() {
	ts := util.ProcessStartTime(8846)
	fmt.Println("pid:8846 start time is :" + ts.String())
}
