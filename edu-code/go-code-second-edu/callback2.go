package main

import (
	"fmt"
	"strings"
	"unicode"
)


func main() {
	// 한글이 처음으로 나타나는 인덱스를 반환
	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c) // c가 한글이면 true를 반환
	}
	fmt.Println(strings.IndexFunc("Hello, 월드", f)) // 7
	fmt.Println(strings.IndexFunc("Hello, world", f)) // -1
}