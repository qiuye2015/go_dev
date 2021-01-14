package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	_, err := Fetch("http://www.baidu.com")
	if err != nil {
		fmt.Printf("Fetch error: %v", err)
	}
	//fmt.Println(string(content))
}
