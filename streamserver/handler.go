package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	//vl := fmt.Sprintf("%s/%s", VIDEO_DIR, vid)
	//video, err := os.Open(vl)
	//if err != nil {
	//	sendErrorResponse(w, http.StatusInternalServerError, "Video open Failed")
	//	return
	//}
	//w.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent(w, r, "", time.Now(), video)
	//defer video.Close()

	// paly video from aliyun oss
	targetUrl := "cyning-video-server.oss-cn-shanghai.aliyuncs.com/videos/" + vid
	http.Redirect(w, r, targetUrl, 301)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SZIE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SZIE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Read file failed: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal service error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file failed: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal service error")
		return
	}
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", VIDEO_DIR, fn), data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Write file error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload success")
}
