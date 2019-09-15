package main

import (
	// "bytes"
	"fmt"
	"os"
	"path"
	

	shell "github.com/memoio/mefs-http-api-go"
)

func main() {
	sh := shell.NewShell("http://212.64.28.207:5001")
	bucketName := "AutoID-B"
	// path
	p := path.Join(os.Getenv("HOME"), "VCV-sample.ddo")
	
	file, err := os.Open(p)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	objectName := path.Base(file.Name())
	// upload file
	_, err = sh.PutObject(file, objectName, bucketName)
	if err != nil {
		fmt.Println("put object err: ", err)
		return
	}

}
