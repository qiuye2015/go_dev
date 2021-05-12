package main

import (
	"fmt"
	"net"
)

var (
	localIPs []string
)

func init() {
	fmt.Println("main init")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("Get local ip failed,err:%v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIPs = append(localIPs, ipnet.IP.String())
			}
		}
	}
	fmt.Println(localIPs)
}
