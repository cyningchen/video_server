package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var (
	EP string
	AK string
	SK string
)

func init() {
	AK = "admin"
	EP = "oss-cn-shanghai.aliyuncs.com"
	SK = "password"
}

func UploadToOss(filename string, path string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("init oss service failed: %v", err)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("get bucket failed: %v", err)
		return false
	}
	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("uploading to bucket failed: %v", err)
		return false
	}
	return true
}
