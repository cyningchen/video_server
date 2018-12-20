package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/user", CreateUser)
	router.POST("/user/:username", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	err := http.ListenAndServe(":8888", r)
	if err != nil{
		fmt.Println(err)
	}
}
