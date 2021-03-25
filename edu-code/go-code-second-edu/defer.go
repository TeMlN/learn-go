// package main

// import (
// 	"fmt"
// )

// func main() {
// 	f1()
// }

// func f1() {
// 	fmt.Println("f1 - start")
// 	defer f2()						//defer 키워드를 사용하면 코드의 위치가 어디던 함수가 끝나기직전 실행된다.
// 	fmt.Println("f1 - end")
// }

// func f2() {
// 	fmt.Printf("f2 - deferred")
// }