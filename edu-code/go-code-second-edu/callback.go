// package main

// import (
// 	"fmt"
// )

// func callback(y int, f func(int, int)) { // main함수에서 선언한대로 y에는 1이 받아지고 func 타입인 f에는 add함수가 담겨진다.
// 	f(y, 2) // add(1, 2)를 호출
// }

// func add(a, b int) {
// 	fmt.Printf("%d + %d = %d\n", a, b, a+b) //1 + 2 = 3
// }

// func main() {
// 	callback(1, add) // 1 + 2 = 3
// }
