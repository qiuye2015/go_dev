package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/qiuye2015/go_dev/balance"
)

func main() {
	// insts := make(, 16)
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port := 8080
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if nil != err {
			fmt.Println("Error", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
