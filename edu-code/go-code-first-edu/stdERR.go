// package main //go의 에러처리 잘 이해못했음으로 다시 공부하기.

// import (
// 	"io"
// 	"os")

// func main() {
// 	myString := ""
// 	arguments := os.Args
// 	if len(arguments) == 1 {
// 		myString = "Please give me one argument!"
// 	} else {
// 		myString = arguments[1]
// 	}

// 	io.WriteString(os.Stdout, "This is Standard output\n")
// 	io.WriteString(os.Stderr, myString)
// 	io.WriteString(os.Stderr, "\n")
// }			 