package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
)

func main() {
	//获取当前网络连接信息
	n1, _ := net.Connections("all") //可填入tcp、udp、tcp4、udp4等等
	fmt.Println("n1:", n1)

	//获取网络读写字节／包的个数
	n2, _ := net.IOCounters(false)
	fmt.Println("n2:", n2)
}
