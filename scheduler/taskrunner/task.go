package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dbops"
)

// 将要删除的video_id 写入到dataChan中
func VideoClearDispatcher(dc dataChan) (err error) {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("video clear dispatcher failed: %v", err)
	}
	if len(res) == 0 {
		return errors.New("All task finished")
	}

	for _, id := range res {
		dc <- id
	}
	return
}

func DeleteVideo(vid string) (err error) {
	err = os.Remove(fmt.Sprintf("%s/%s", VIDEO_PATH, vid))
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Delete video file failed: %v", err)
		return
	}
	return
}

func VideoClearExecutor(dc dataChan) (err error) {
	errMap := &sync.Map{}
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err = DeleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err = dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = key.(error)
		if err != nil {
			return false
		}
		return true
	})
	return
}
