package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 设置split
func TestScannerSplit() {
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 设置split方法模式
	// bufio.Scanner 实例的 split为 splitFunc的实例
	scanner.Split(bufio.ScanWords)
	count := 0
	// 相当于 iterator的 next
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
}

// scanner按行读取文件
func TestScannerText() {
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("it is first line\nit is second line"))
	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner false", err)
	}
}

func main() {
	TestScannerText()
}
