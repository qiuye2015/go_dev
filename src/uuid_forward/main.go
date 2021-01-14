package main

import (
	"fmt"
	"strings"
)

const (
	path = "D:\\code\\go_dev\\src\\uuid.txt"
	//respath  = "D:\\code\\go_dev\\src\\result.txt"
	respath = "D:\\code\\go_dev\\src\\result2.txt"
	url_pre = "http://10.13.3.36:8899/s/trace?uuid="
)

func main() {
	resMap := ReadFromFile(path)
	var buffer strings.Builder

	for uuid, _ := range *resMap {
		ret := GetPage(url_pre + uuid)
		fwd := ret.Result.Data.Indexed.Source
		buffer.WriteString(uuid + " " + fwd.Surl + "\n")
		fmt.Println(uuid, fwd.Surl, fwd.Media, fwd.Muid)
	}
	WriteToFile(respath, buffer.String())

}
