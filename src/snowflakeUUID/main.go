package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	startTime   int64 = 1620458281000           //如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	workerBits  uint8 = 10                      //10 bit 机器标识
	numberBits  uint8 = 12                      //12bit自增ID
	workerMax   int64 = -1 ^ (-1 << workerBits) //最多1023台机器 2^10-1
	numberMax   int64 = -1 ^ (-1 << numberBits) //最大自增id是4096 2^1201
	timeShift   uint8 = workerBits + numberBits //22bit前是毫秒时间戳
	numberShift uint8 = workerBits              //10bit前是自增id
	//workerShift uint8 = numberBits              //12bit前是机器标识

)

type Worker struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerId      int64
	number        int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		lastTimestamp: 0,
		workerId:      workerId,
		number:        0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.lastTimestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
			w.number = 0          //BY FJP
			w.lastTimestamp = now //By FJP
		}
	} else {
		w.number = 0
		w.lastTimestamp = now
		if now < w.lastTimestamp { //时间回拨
			log.Println("Clock moved backwards.  last time is ", w.lastTimestamp, "current time is ", now)
			//解决时间回拨的问题
			//now = w.lastTimestamp
		}
	}
	//ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	ID := int64((now-startTime)<<timeShift | (w.number << numberShift) | (w.workerId))
	return ID
}

func main() {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(node.GetId())
		//fmt.Println(node.GetId() & 0b1111111111) //workID
		//fmt.Println(node.GetId() & 0b1111111111110000000000) //自增ID
		//fmt.Println(node.GetId() & 0b01111111111111111111111111111111110000000000000000000000) //时间戳
		//time.Sleep(time.Second)
		//time.Sleep(time.Millisecond * 500)
		//time.Sleep(time.Microsecond)
		//time.Sleep(time.Nanosecond)
	}
}
