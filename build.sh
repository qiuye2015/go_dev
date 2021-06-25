#!/bin/bash
#########################################################################
# File Name: build.sh
# Author: fjp
# mail: fjp@xxx.com
# Created Time: Thu Apr 30 11:01:38 2020
#########################################################################

src_file=$1
if [ ! $src_file ];then
	echo Usage: `basename $0` src_directory
	exit
fi
go build -gcflags=all="-N -l" $1
## 必须这样编译，才能用gdb打印出变量，第二个是小写的L，不是大写的i
#go build  -o stock_monitor src/prometheus/main.go
