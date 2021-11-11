package main

import (
	"fmt"
	"unicode/utf8"
)

// 判断Unicode编码效果

// 判断是否符合utf-8编码：Valid, ValidRune, ValidString
// 判断rune所占字节数: RuneLen
// 判断字节串或者字符串的rune数: RuneCount, RuneCountInString
// 编码和解码到rune: EncodeRune, DecodeRune, DecodeRuneInString, DecodeLastRune
// 是否为完整rune: FullRune, FullRuneInString
// 是否为rune第一个字节: RuneStart

func main() {
	word := []byte("界")
	fmt.Println(utf8.Valid(word))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))
	fmt.Println(utf8.RuneLen('界'))
	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界"))
	p := make([]byte, 3)
	utf8.EncodeRune(p, '界')
	fmt.Println(p)
	r, _ := utf8.DecodeRune(p)
	fmt.Printf("%c\n", r)
	// 只能解出一个字符
	c, _ := utf8.DecodeRuneInString("你好")
	fmt.Printf("%c", c)
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))
	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}
