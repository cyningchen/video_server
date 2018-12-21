package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MiddWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddwareHandler(r *httprouter.Router, bucket int) http.Handler {
	m := MiddWareHandler{}
	m.r = r
	m.l = NewConnLimiter(bucket)
	return m
}

func (m MiddWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddwareHandler(r, 2)
	if err := http.ListenAndServe(":9000", mh); err != nil {
		fmt.Println(err)
	}
}
