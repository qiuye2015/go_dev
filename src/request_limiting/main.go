package main

import (
	"log"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// 基于 IP 地址的限速器
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)
	if err := http.ListenAndServe(":8888", limitMiddleware(mux)); err != nil {
		//if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatal("unable to start server: %s", err.Error())
	}
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all get"))
}

// 令牌桶限速算法 x/time/rate
// 它实现了一个大小为 b 的「令牌桶」，初始化时是满的，并且以每秒 r 个令牌的速率重新填充。
// 非正式地，在任意足够长的时间间隔中，限速器将速率限制在每秒 r 个令牌，最大突发事件为 b 个
// IPRateLimiter
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter .
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	i.mu.Unlock()
	return limiter
}

var limiter = NewIPRateLimiter(1, 5)
var count int

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			log.Print("invalid url")
			return
		}
		log.Print("----", count)
		count++
		next.ServeHTTP(w, r)
	})
}
