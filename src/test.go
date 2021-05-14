package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var count int32

func main() {
	TestDoDupSuppress()
	//
	//loadGroup := &singleflight.Group{}
	////res := []<-chan singleflight.Result{}
	//
	//for i := 0; i < 10; i++ {
	//	//key := "hello" + strconv.Itoa(i)
	//	key := "hello"
	//	go func() {
	//		resVec, err, shared := loadGroup.Do(key, GetUserVec)
	//		if err != nil {
	//			fmt.Println("loadGroup get userVec erro")
	//		}
	//		fmt.Println(resVec, shared)
	//	}()
	//}
	time.Sleep(time.Second)

	//for i := 0; i < 10; i++ {
	//	dat := <-res[i]
	//	fmt.Println(dat.Val)
	//}
}

func GetUserVec() (interface{}, error) {
	a := int(rand.Int31() % 1000)
	time.Sleep(time.Millisecond * time.Duration(a))
	iter := atomic.AddInt32(&count, 1)
	fmt.Println("GetUserVec----", iter, a)
	return iter, nil
}

func TestDoDupSuppress() {
	var g singleflight.Group
	c := make(chan string)
	var calls int32
	fn := func() (interface{}, error) {
		atomic.AddInt32(&calls, 1)
		return <-c, nil
	}

	const n = 10
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() { // n个协程同时调用了g.Do，fn中的逻辑只会被一个协程执行
			v, err, shard := g.Do("key", fn)
			if err != nil {
				fmt.Println("Do error: %v", err, shard)
			}
			if v.(string) != "bar" {
				fmt.Println("got %q; want %q", v, "bar", shard)
			}
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond) // let goroutines above block
	c <- "bar"
	wg.Wait()
	if got := atomic.LoadInt32(&calls); got != 1 {
		fmt.Println("number of calls = %d; want 1", got)
	}
}
