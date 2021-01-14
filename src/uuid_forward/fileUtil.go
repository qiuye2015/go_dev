package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFromFile(path string) *map[string]struct{} {
	resMap := make(map[string]struct{}, 0)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	buffer := bufio.NewReader(f)

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err != io.EOF || line == "" {
				break
			}
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			resMap[line] = struct{}{}
		}
	}
	return &resMap
}

func WriteToFile(path, content string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
	}
	defer file.Close()

	file.WriteString(content)
}
