package main

import (
	"io"
	"os"
	"strings"
)

// Copy 将 src 复制到 dst，直到在 src 上到达 EOF 或发生错误。
// 它返回复制的字节数，如果有错误的话，还会返回在复制时遇到的第一个错误。

// 成功的 Copy 返回 err == nil，而非 err == EOF。
// 由于 Copy 被定义为从 src 读取直到 EOF 为止，因此它不会将来自 Read 的 EOF 当做错误来报告。

// 若 dst 实现了 ReaderFrom 接口，其复制操作可通过调用 dst.ReadFrom(src) 实现。
// 此外，若 src 实现了 WriterTo 接口，其复制操作可通过调用 src.WriteTo(dst) 实现。
func TestCopy() {
	io.Copy(os.Stdout, strings.NewReader("copy is running\n"))
	io.CopyN(os.Stdout, strings.NewReader("copy is running\n"), 8)
}

func main() {
	TestCopy()
}
