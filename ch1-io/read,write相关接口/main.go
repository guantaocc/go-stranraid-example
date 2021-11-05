package main

// 基础的io.reader io.readAt接口示例
// ReaderFrom, WriterTo

import (
	"fmt"
	"io"
	"os"
	"strings"
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

// ReadFrom接口实现
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

// 拼接 byte 字节
func readAllFunc(file *os.File) (b []byte, err error) {
	p := make([]byte, 64)
	for {
		n, err := file.Read(p)
		if n > 0 {
			bytes := p[:n]
			b = append(b, bytes...)
		}
		if err == io.EOF {
			return b, nil
		}
		if err != io.EOF && err != nil {
			return b, err
		}
	}
}

// writeTo 接口示例: 将全部byte写入
func TestWriteTo() {
	str := strings.NewReader("this is io writeto")
	str.WriteTo((os.Stdout))
}

func main() {
	// 从标准输入中读数据，比如os.stdin, scan
	// s := "hello world 中文网"
	fd, _ := os.OpenFile("../test.txt", os.O_RDWR, os.ModePerm)
	defer fd.Close()
	b, _ := readAllFunc(fd)
	fmt.Println(string(b))

	TestWriteTo()
}
