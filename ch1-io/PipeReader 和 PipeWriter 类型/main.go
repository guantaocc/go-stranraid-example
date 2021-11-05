package main

import (
	"errors"
	"fmt"
	"io"
	"time"
	"unicode/utf8"
)

// PipeReader（一个没有任何导出字段的 struct）是管道的读取端。它实现了 io.Reader 和 io.Closer 接口。

func PipeRead(reader *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接口端开始阻塞5秒钟...")
		time.Sleep(5 * time.Second)
		fmt.Println("接收端开始接受")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		r, _ := utf8.DecodeRune(buf)
		fmt.Printf("收到字节: %d\n buf内容: %c\n", n, r)
	}
}

func PipeWriter(writer *io.PipeWriter) {
	data := "这是pipe输入端"
	// 单个字符发送
	for _, r := range data {
		// 计算rune的长度
		p := make([]byte, 4)
		len := utf8.EncodeRune(p, r)
		writer.Write(p[:len])
	}
	writer.CloseWithError(errors.New("写入段已关闭"))
}

func main() {
	reader, writer := io.Pipe()
	go PipeRead((reader))
	go PipeWriter(writer)
	time.Sleep(30 * time.Second)
}
