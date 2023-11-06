package main

import (
	"fmt"
	"net"
)

func main() {
	//
	ipv4Macs := GetLocalIpv4Macs()

	var macs []string
	var ips []string

	for k, v := range ipv4Macs {
		macs = append(macs, k)
		ips = append(ips, v)
	}

	fmt.Printf("macs: %v\n", macs)
	fmt.Printf("ips: %v\n", ips)
}

func GetLocalIpv4Macs() map[string]string {
	mp := make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		return mp
	}

	for _, inter := range interfaces {
		addrs, _ := inter.Addrs()
		for _, address := range addrs {
			if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Printf("ip: %v, mac: %v\n", ipNet.IP.String(), inter.HardwareAddr.String())
					mp[inter.HardwareAddr.String()] = ipNet.IP.String()
				}

			}
		}
	}

	return mp
}
