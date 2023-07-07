package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("err:%s \n", err.Error())
		return
	}

	fmt.Printf("interfaces:%v \n", interfaces)
	for _, inter := range interfaces {
		fmt.Printf("=================================================\n")
		fmt.Printf("inter.HardwareAddr. mac:%s\n", inter.HardwareAddr)
		addrs, _ := inter.Addrs()
		fmt.Printf("inter.Addrs(). addrs:%v \n", addrs)
		for _, address := range addrs {
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Printf("ip:%v\n", ipNet.IP.String())
				}

			}
		}
	}
}
