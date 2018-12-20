package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	io.WriteString(w, "create user handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("username")
	io.WriteString(w,uname)
}