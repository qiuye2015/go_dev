package main

import (
	"fmt"
	"io"
)

type fake struct{ io.Writer }

func fred(logger io.Writer) {
	fmt.Printf("%#v\n", logger)
	if logger != nil {
		logger.Write([]byte("..."))
	}
}

func main() {
	var a *string = nil
	var b interface{} = nil
	fmt.Println("a == nil:", a == nil) // true
	fmt.Println("b == nil:", b == nil) // true
	fmt.Println("a == b:", a == b)     // false (尽管a和b的值都为nil)
}
