module github.com/qiuye2015/go_dev/GrpcCode

go 1.14

require (
	github.com/qiuye2015/go_dev/GrpcCode/message v0.0.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0 // indirect
)

replace github.com/qiuye2015/go_dev/GrpcCode/message => ./message
