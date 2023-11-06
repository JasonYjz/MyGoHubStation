package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var ipmacs []string

	for _, inter := range interfaces {
		fmt.Println(inter.Name)
		addrs, _ := inter.Addrs()
		for _, address := range addrs {

			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ipmacs = append(ipmacs, ipNet.IP.String(), inter.HardwareAddr.String())
				}

			}
		}
	}

	fmt.Println(ipmacs)
}
