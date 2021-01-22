package main

import (
	"context"
	"filetuil"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
	"time"
)

var serverMux = http.NewServeMux()

type GoServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

func NewGoServer(addr string) *GoServer {
	var goServer GoServer
	if addr != "" {
		goServer.server = &http.Server{
			Addr:         addr,
			Handler:      serverMux,
			ReadTimeout:  time.Second * 1, //设置3秒的读超时
			WriteTimeout: time.Second * 3, //设置3秒的写超时
		}
		goServer.wg.Add(1)
		go func() {
			if err := goServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Println("server listen error", addr, err)
			}
			goServer.wg.Done()
		}()
	}
	return &goServer
}

func (s *GoServer) GoShutDown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}
	if err := s.server.Shutdown(ctx); err != nil { //Shutdown-->Close
		return err
	}
	s.wg.Wait()
	return nil
}

//定义两个对外注册函数
func GoHandle(pattern string, handler http.Handler) {
	serverMux.Handle(pattern, handler)
}

//func GoHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
//	serverMux.HandleFunc(pattern, handler)
//}

func main() {
	flag.Parse()
	metricServer := NewGoServer(*addr)
	defer metricServer.GoShutDown(context.Background())
	log.Println("starting server: ", *addr)
	watchStock()
	//TODO:
	// 等待信号
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		sig := <-sigChan
		log.Println("recv signal exit", sig)
		break
	}
}
func init() {
	GoHandle("/metrics", promhttp.Handler())
}

var (
	stockPro = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fjp_stock",
		Help: "fjp_stock",
	}, []string{"company"})
	addr = flag.String("listen-address", ":9999", "The address to listen on for HTTP requests.")
	path = flag.String("watch_file", "./text", "The path to monitor")
)

func watchStock() {
	actionMap := make(map[string]func(s string) error)
	if err := load(*path); err != nil {
		log.Println(err)
	}
	actionMap[*path] = load
	if len(actionMap) > 0 {
		go filetuil.FileWatcher(actionMap)
	}
}

func load(s string) error {
	tmp, err := filetuil.ReadFromFile(s)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_resMap := make(map[string]float64, 0)
	for k, _ := range tmp {

		Vec := strings.Split(k, ". ")
		if len(Vec) != 2 {
			continue
		}
		kv := strings.Split(Vec[1], "：")
		if len(kv) != 2 {
			continue
		}
		marketValue := strings.TrimRight(kv[1], "亿美元")
		marketValueNum, _ := strconv.ParseFloat(marketValue, 64)
		_resMap[kv[0]] = marketValueNum
		//log.Println(kv[0], marketValueNum)
		stockPro.With(prometheus.Labels{"company": kv[0]}).Set(marketValueNum)
	}
	return nil
}
