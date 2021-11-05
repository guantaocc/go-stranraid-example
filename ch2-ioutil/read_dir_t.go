package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func TestReadDir() {
	// 读取目录
	path := os.Args[1]
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, info := range fileInfos {
		fmt.Println(info.Name())
	}
}

func main() {
	TestReadDir()
}
