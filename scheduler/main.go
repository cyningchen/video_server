package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/taskrunner"
)

func main() {
	go taskrunner.Start()
	r := RegisterHandlers()
	if err := http.ListenAndServe("9001", r); err != nil {
		fmt.Println("listen and serve failed, ", err)
	}

}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", videoDelRecHandler)
	return router
}
