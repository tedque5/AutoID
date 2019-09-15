package main

import (
	// "bytes"
	"fmt"
	"os"
	"path"

	shell "github.com/memoio/mefs-http-api-go"
)

func GetCurrentDirectory() string {
   curDir,err := filepath.Abs(filepath.Dir(os.Args[0]))  
   if err != nil {
      log.Fatal(err)
   }
   return strings.Replace(curDir, "\\", "/", -1) 
}

func main() {
	sh := shell.NewShell("http://212.64.28.207:5001")
	bucketName := "AutoID-B"
	// 文件路径
	//p := path.Join(os.Getenv("HOME"), "VCV-sample.ddo")
	curDir = GetCurrentDirectory()
	p := path.Join(curDir, "VCV-sample.ddo")

	file, err := os.Open(p)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}

	objectName := path.Base(file.Name())
	// 上传文件
	_, err = sh.PutObject(file, objectName, bucketName)
	if err != nil {
		fmt.Println("put object err: ", err)
		return
	}

}
