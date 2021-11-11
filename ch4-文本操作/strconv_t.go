package main

import (
	"fmt"
	"strconv"
)

// strconv: 提供字符串和相关基础数据类型的转换
// 基本数据类型: bool, int(有无符号), 二/八/十六进制, 浮点型
// ErrRange 和 ErrSyntax。其中，ErrRange 表示值超过了类型能表示的最大范围，比如将 "128" 转为 int8 就会返回这个错误；ErrSyntax // 表示语法错误，比如将 "" 转为 int 类型会返回这个错误。

// 字符串转换为整型：ParseInt, ParseUint, Atoi
// 整形转换字符串: FormatUint, FormatInt, Itoa
// 字符串和bool转换：ParseBool, FormatBool, AppendBool
// 字符串和浮点数转换：ParseFloat, FormatFloat, AppendFloat

func main() {
	n, err := strconv.ParseInt("128", 10, 8)
	fmt.Println(n, err)
	// 127, ErrRange类型
	_, ok := err.(*strconv.NumError)
	fmt.Print(ok)
}
