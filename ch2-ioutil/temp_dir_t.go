package main

import (
	"io/ioutil"
	"os"
)

// 测试ioutil创建临时文件的功能

func TestTmp() {
	f1, err := ioutil.TempFile("", "test")
	if err != nil {
		panic(err)
	}
	f1.WriteString("这是临时文件")

	defer func() {
		f1.Close()
		os.Remove(f1.Name())
	}()
}

func main() {
	TestTmp()
}
