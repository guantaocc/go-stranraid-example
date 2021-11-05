package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func TestMultiReader() {
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	reader := io.MultiReader(readers...)

	buf := make([]byte, 128)

	// 有几个reader分别调用之后才会返回io.EOF
	for _, err := reader.Read(buf); err != io.EOF; _, err = reader.Read(buf) {
		fmt.Println(string(buf))
	}
}

func TestMultiWriter() {
	file, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	// write一次即可
	writer.Write([]byte("Go语言中文网"))
}

func main() {
	TestMultiReader()
	TestMultiWriter()
}
