package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func getWorkId(Host string) int {
	workerId := 0
	ip := strings.Split(Host, ".")
	if len(ip) == 4 {
		for _, value := range ip {
			ipValue, _ := strconv.Atoi(value)
			workerId += ipValue & 0xFF
		}
	} else {
		workerId = rand.Intn(1024)
	}
	return workerId
}

func main() {
	host_list := []string{
		"255.255.255.255",
	}
	for _, v := range host_list {
		fmt.Println(getWorkId(v), v)
	}

}
