package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

func main() {
	//printCpuTimes()
	getCpuInfo()
}

// 获取Cpu方面的信息
func printCpuTimes() {
	res, err := cpu.Times(true) // false是展示全部总和 true是分布展示
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(5 * time.Second, true)
		fmt.Println("cpu percent:%v", percent)
	}
}


