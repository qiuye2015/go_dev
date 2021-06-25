package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				//ch <- struct{}{}
				fmt.Println("结束go")
				return
			default:
				fmt.Println("default...")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	//<-ch
	time.Sleep(10 * time.Second)
	fmt.Println("结束")
}
