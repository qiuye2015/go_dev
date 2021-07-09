package redlock

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test_Main(t *testing.T) {
	log.SetFlags(log.Ltime | log.Llongfile)
	//rc1 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7001"})
	//rc2 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7002"})
	//rc3 := redis.NewClient(&redis.Options{Addr: "0.0.0.0:7003"})
	//
	//dlm := NewDLM([]*redis.Client{rc1, rc2, rc3}, 10*time.Second, 2*time.Second)
	rc := redis.NewClient(&redis.Options{Addr: "10.22.5.25:16379"})
	dlm := NewDLM([]*redis.Client{rc}, 10*time.Second, 2*time.Second)

	withLockOnly(dlm)

	withLockAndUnlock(dlm)
}

func withLockAndUnlock(dlm *DLM) {
	ctx := context.Background()
	locker := dlm.NewLocker("this-is-a-key-002")

	if err := locker.Lock(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	if err := locker.Unlock(ctx); err != nil {
		log.Fatal("[main] Failed when unlocking, err:", err)
	}

	log.Println("[main] Done")
}

func withLockOnly(dlm *DLM) {
	ctx := context.Background()
	locker := dlm.NewLocker("this-is-a-key-002")

	if err := locker.Lock(ctx); err != nil {
		log.Fatal("[main] Failed when locking, err:", err)
	}

	// Perform operation.
	someOperation()

	// Don't unlock

	log.Println("[main] Done")
}

func someOperation() {
	log.Println("[someOperation] Process has been started")
	time.Sleep(1 * time.Second)
	log.Println("[someOperation] Process has been finished")
}

func Test_generateRandomString(t *testing.T) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune,
		time.Now().Unix()%64)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	fmt.Println(string(b))
}
