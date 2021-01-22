package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//httpServerV1()
	//httpServerV2()
	httpServerV3()
}

func httpServerV1() {
	http.HandleFunc("/v1", sayHello)

	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Starting server ..."))
}

//------------------------------------------------
func httpServerV2() {
	mux := http.NewServeMux()
	mux.Handle("/v2", &myHandler{})
	mux.HandleFunc("/v2_2", sayHello)

	log.Println("Starting v2 server ...")
	log.Fatal(http.ListenAndServe(":9999", mux)) //nil-->mux
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is ServerHTTP"))
}

//------------------------------------------------
func httpServerV3() {
	mux := http.NewServeMux()
	mux.Handle("/v3", &myHandler{})

	server := &http.Server{
		Addr:         ":9999",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	log.Println("Starting v3 httpserver")
	log.Fatal(server.ListenAndServe())
}
