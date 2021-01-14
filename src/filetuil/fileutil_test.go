package filetuil

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

var resMap *map[string]struct{}

const (
	path = "test.txt"
)

func TestLoad(t *testing.T) {
	_tmpMap := make(map[string]struct{}, 0)
	resMap = &_tmpMap

	actionsMap := make(map[string]func(s string) error)
	if err := load(path); err != nil {
		fmt.Println(err)
	}
	fmt.Println("add load")
	actionsMap[path] = load
	if len(actionsMap) > 0 {
		go FileWatcher(actionsMap)
	}
	fmt.Println("waiting")
	time.Sleep(time.Second * 130)
}

func load(s string) error {
	fmt.Println("load starting...")
	_resMap, err := ReadFromFile(s)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if len(_resMap) > 0 {
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&resMap)),
			unsafe.Pointer(&_resMap))

		fmt.Println("len:", len(*resMap))
		for k, _ := range *resMap {
			fmt.Println(k)
		}
	}
	return nil
}
