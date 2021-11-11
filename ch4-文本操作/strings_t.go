package main

import (
	"fmt"
	"strings"
)

// 测试字符串相关操作
// 比较 Compare, EqualFold(忽略大小写)
// 存在或子串 Contains ContainsAny ContainsRune
// 字符串出现次数 Count
// 字符串分割 Fields FieldsFunc Split SplitAfter SplitN SplitAfterN
// 字符串是否由前后缀 HasPrefix HasSuffix
// 子字符串出现的位置 Index, IndexByte, IndexAny, IndexFunc, IndexRue. LastIndex ...
// 拼接字符串 Join
// 字符串重复 Repeat
// 字符替换 Map Replace ReplaceAll
// 大小写转换 ToLower ToUpper ToUpperSpecial
// 标题处理，转换驼峰等等 Title ToTitle ToTitleSpecial
// 修剪字符 Trim TrimLeft TrimRight TrimPrefix TrimSuffix TrimSpace TrimFunc TrimLeftFunc TrimRightFunc
// Replacer 类型 成对替换
// Reader => io.Reader
// Builder => 实例化Writer相关操作

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
	b := &strings.Builder{}
	_, _ = b.WriteString("hello world")
	fmt.Printf("%s", b.String())
}
