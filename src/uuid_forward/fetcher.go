package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPage(url string) *UuidForward {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http request failed", err)
		return nil
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, 1024*10))
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("http get msg failed", err)
		return nil
	}
	var uuidFwd UuidForward
	err = json.Unmarshal(buf.Bytes(), &uuidFwd)
	if err != nil {
		fmt.Println("json Unmarsshal failed", err)
	}
	return &uuidFwd
}
