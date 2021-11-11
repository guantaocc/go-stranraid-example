package main

import "os"

func TestFileSeek() {
	// 由于seek产生的空洞文件
	file, _ := os.Create("test.txt")
	defer file.Close()
	file.Seek(10, os.SEEK_SET)
	file.WriteString("hello world")
}

func main() {

}
