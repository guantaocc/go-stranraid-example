package main

import (
	"bufio"
	"fmt"
	"strings"
)

// bufio的 Reader 和 Writer结构体实现了 io.Reader 和 io.Writer并提供缓冲的结构

// type Reader struct {
// 	buf          []byte        // 缓存
// 	rd           io.Reader    // 底层的io.Reader
// 	// r:从buf中读走的字节（偏移）；w:buf中填充内容的偏移；
// 	// w - r 是buf中可被读的长度（缓存数据的大小），也是Buffered()方法的返回值
// 	r, w         int
// 	err          error        // 读过程中遇到的错误
// 	lastByte     int        // 最后一次读到的字节（ReadByte/UnreadByte)
// 	lastRuneSize int        // 最后一次读到的Rune的大小 (ReadRune/UnreadRune)
// }

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com\nit is next line"), 100)
	line, isPrefix, err := reader.ReadLine()
	fmt.Printf("line:%s, %t, %v\n", line, isPrefix, err)
	line, isPrefix, err = reader.ReadLine()
	fmt.Printf("line:%s, %t, %v\n", line, isPrefix, err)
}
