package main

import (
	"bytes"
	"fmt"
)

// bufio.Reader/Writer 分别实现了io.ByteReader 和 io.ByteWriter
// bytes.Buffer 同时实现了 io.ByteReader 和 io.ByteWriter
// bytes.Reader 实现了 io.ByteReader
// strings.Reader 实现了 io.ByteReader

func TestByteReader() {
	buffer := &bytes.Buffer{}
	buffer.WriteByte('b')
	fmt.Println("写入字节成功")
	ch, err := buffer.ReadByte()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("读取字节%c", ch)
}

func main() {
	TestByteReader()
}
