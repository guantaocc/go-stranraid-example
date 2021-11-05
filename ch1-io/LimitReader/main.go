package main

import (
	"fmt"
	"io"
	"strings"
)

// 从 R 读取但将返回的数据量限制为 N 字节。每调用一次 Read 都将更新 N 来反应新的剩余数量。

func main() {
	str := "this is limitReader test"
	reader := strings.NewReader(str)
	// 最多读8个字节
	limitReader := &io.LimitedReader{R: reader, N: 8}
	buf := make([]byte, 2)
	for limitReader.N > 0 {
		limitReader.Read(buf)
		fmt.Printf("%s", buf)
	}
}
