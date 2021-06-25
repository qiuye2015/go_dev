# go环境安装
export GOPATH=`pwd`
#mkdir bin
#mkdir pkg
#mkdir -p src/golang.org/x
#cd src/golang.org/x
#git clone https://github.com/golang/tools.git
#wget https://github.com/golang/tools/archive/gopls/v0.6.4.tar.gz
#tar zxf v0.6.4.tar.gz
#mv tools-gopls* tools
#cd -
go get -u -v github.com/nsf/gocode
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/josharian/impl
#go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v github.com/sqs/goreturns
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/cweill/gotests/
go get -u -v github.com/go-delve/delve/cmd/dlv
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v golang.org/x/tools/cmd/godoc

