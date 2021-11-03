package main

import (
	"fmt"
	"io"
	"os"
)

// 基本IO接口

func readFrom(r io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := r.Read(p)
	if n > 0 {
		return p[:n], err
	}
	return p, err
}

func readAll(file *os.File) (b []byte, err error) {
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(fileInfo.Size())
	var p = make([]byte, fileInfo.Size())
	_, err = file.Read(p)
	return p, err
}

func main() {
	// 从标准输入中读数据，比如os.stdin, scan
	// s := "hello world 中文网"
	fd, _ := os.OpenFile("../test.txt", os.O_RDWR, os.ModePerm)
	defer fd.Close()
	b, _ := readAll(fd)
	fmt.Println(string(b))
}
