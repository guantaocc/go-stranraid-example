package main

import (
	"fmt"
	"io"
	"strings"
)

// 主要介绍io.seeker
// Seek 设置下一次 Read 或 Write 的偏移量为 offset，
// 它的解释取决于 whence：
// 0 表示相对于文件的起始处，1 表示相对于当前的偏移，而 2 表示相对于其结尾处。
// Seek 返回新的偏移量和一个错误，如果有的话。

func TestSeek() {
	s := strings.NewReader("seek 中文及示例")
	// 从尾部偏移6个字节
	s.Seek(-12, io.SeekEnd)
	r, _, _ := s.ReadRune()
	fmt.Printf("%c\n", r)
}

func main() {
	TestSeek()
}
