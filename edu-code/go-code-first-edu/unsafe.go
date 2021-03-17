package main

import (
	"fmt"
	"unsafe"
)


func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))	 //포인터p1을 unsafe.Pointer()함수로 p2로 복사 (int32)타입으로

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)                //32비트 포인터로는 64비트 정수를 담을 수 없다.
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}