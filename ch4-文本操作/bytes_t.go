package main

import (
	"bytes"
	"fmt"
)

// 字节数组的相关操作
// 存在或子字节 Contains
// 出现次数 Count
// 类型转换 Runes
// Reader类型 NewReader
// Buffer类型 NewBufferString NewBuffer bytes.Buffer{}
// Buffer的相关函数 Truncate ReadBytes ReadString

// 字节数组转换Rune数组
func TestRunes() {
	b := []byte("你好，世界")
	for k, v := range b {
		fmt.Printf("%d:%s |", k, string(v))
	}
	r := bytes.Runes(b)
	for k, v := range r {
		fmt.Printf("%d:%s|", k, string(v))
	}
}

func main() {
	TestRunes()
}
