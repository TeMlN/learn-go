// package main

// import (
// 	"fmt"
// )

// func f() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Printf("%d ", i)	//출력 결과를 보면 4 3 2 1 0 이 출력이 돼는데 그 이유는 defer로 인해 출력되지 못하고 stack에 쌓였기 때문이다.
// 	}
// }

// func main() {
// 	f()
// }