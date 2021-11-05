package main

import (
	"bytes"
	"fmt"
)

// ByteScanner: 将上一次 ReadByte 的字节还原

func main() {
	buffer := bytes.NewBuffer([]byte{'a', 'b'})
	r, _ := buffer.ReadByte()
	fmt.Print(r)
	buffer.UnreadByte()
	re, _ := buffer.ReadByte()
	fmt.Print(re == r)
}
