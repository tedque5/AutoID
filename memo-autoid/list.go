package main

import (
	"fmt"

	shell "github.com/memoio/mefs-http-api-go"
)

func main() {
	sh := shell.NewShell("http://212.64.28.207:5001")
	bucketName := "AutoID-B"
	
	// buckets 列表
	bks, err := sh.ListBuckets()
	fmt.Println(bks,err)

	// 文件列表
	obs, err := sh.ListObjects(bucketName)
	if err != nil {
		fmt.Println("list object err: ", err)
		return
	}
	fmt.Println(obs,err)
}