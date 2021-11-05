package main

import (
	"fmt"
	"io"
)

// 文件 (os.File)、归档（压缩包）、数据库连接、Socket 等需要手动关闭的资源都实现了 Closer 接口。

type CustomCloser struct {
	io.Closer
}

func (c *CustomCloser) Close() error {
	fmt.Println("is CustomCloser closed")
	return nil
}

func TestCloser() {
	c := &CustomCloser{}
	c.Close()
}

func main() {
	TestCloser()
}
