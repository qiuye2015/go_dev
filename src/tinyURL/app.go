package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	Middlewares *Middleware
	Conf        Config
	S           StorageInf
}

func (a *App) InitApp() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	a.Router = mux.NewRouter()
	a.Middlewares = &Middleware{}
	a.Conf.InitConfig()
	a.S = NewRedisCli(a.Conf.Addr, a.Conf.Db)
}

func (a *App) Run(addr string) {
	log.Printf("Http Server Run on [%v]", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// InitRouter
/*
============================================================
http://localhost:9999/api/shorten
Content-Type application/json
{
    "url":"http://www.baidu.com",
    "expiration_in_minutes": 30
}
重定向为相对地址,要想使用绝对地址,url必须是全的,加上http的
============================================================
http://localhost:9999/api/info?shortlink=1
============================================================
http://localhost:9999/api/1
============================================================
*/
func (a *App) InitRouter() {
	//a.Router.HandleFunc("/hello", hello)
	m := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)
	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortLink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortInfo)).Methods("GET")
	a.Router.Handle("/api/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(a.redirect)).Methods("GET")
}

//func hello(w http.ResponseWriter, r *http.Request) {
//	//log.Printf("Enter hello [%v]: %s ", r.Method, r.URL)
//	w.Write([]byte("hellp http server"))
//}

type shortenReq struct {
	//长地址
	URL string `json:"url" validate:"required"`
	//过期时间
	Expiration int64 `json:"expiration_in_minutes" validate:"min=0"`
}
type shortenResp struct {
	ShortLink string `json:"short_link"`
}

//createShortLink 创建短地址函数
func (a *App) createShortLink(w http.ResponseWriter, r *http.Request) {
	//a.S.Shorten()
	var req shortenReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseWithError(w, StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("parse parameters failed %v", err),
		})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		responseWithError(w, StatusError{
			Code: http.StatusBadRequest,
			Err:  fmt.Errorf("validate parameters failed %v", err),
		})
		return
	}
	defer r.Body.Close()
	//fmt.Printf("%+v\n", req)
	s, err := a.S.Shorten(req.URL, req.Expiration)
	if err != nil {
		responseWithError(w, err)
	} else {
		responseWithJson(w, http.StatusCreated, shortenResp{
			ShortLink: s,
		})
	}
}

func (a *App) getShortInfo(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("shortlink")
	//fmt.Printf("%s\n", s)
	d, err := a.S.ShortLinkInfo(s)
	if err != nil {
		responseWithError(w, err)
	} else {
		responseWithJson(w, http.StatusOK, d)
	}

}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//log.Printf("%s\n", vars["shortlink"])
	u, err := a.S.UnShorten(vars["shortlink"])
	if err != nil {
		log.Printf("%v", err)
		responseWithError(w, err)
	} else {
		log.Printf("redirect:%v", u)
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}
}
func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.WriteHeader(code) //写入错误码
	w.Header().Set("Context-Type", "application/json")
	w.Write(resp)
}

//返回给客户端json信息
func responseWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error: //如果是自定义的类型错误
		log.Printf("HTTP %d - %s", e.Status(), e.Error())
		responseWithJson(w, e.Status(), e)
	default: //默认的Interal错误
		code := http.StatusInternalServerError
		responseWithJson(w, code, http.StatusText(code))
	}
}
