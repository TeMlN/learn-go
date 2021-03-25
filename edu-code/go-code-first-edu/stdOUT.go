// package main

// import (
// 	"io"
// 	"os"
// )

// func main() {
// 	myString := "" 									//화면에 출력할 변수
// 	arguments := os.Args
// 	if len(arguments) == 1 {
// 		myString = "Please give me one argument!"
// 	} else {
// 		myString= arguments[1]
// 	}

// 	io.WriteString(os.Stdout, myString)			    //fmt.Print()함수와 똑같이 작용하지만 두개의 매개변수만 받는다.
// 	io.WriteString(os.Stdout, "\n")					//io.WriteString()함수는 두개의 매개변수를 받고. 2번째 매개변수으 내용을 1번째로 보내서 화면에 출력한다.
// } 