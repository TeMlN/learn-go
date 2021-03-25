// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	array := [...]int{0, 1, -2, 3, 4}
// 	pointer := &array[0]																//array[0]에 대한 메모리 주소를 가리킴.
// 	fmt.Print(*pointer, " ")
// 	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])			//pointer를 unsafe.Pointer()함수로 변환 시키고 다시 uintptr로 변환한다 그리고 다시 memoryAddress에 저장.

// 																						//여기서 uintptr이란, uint와 크기가 동일하며 포인터를 저장할 때 사용

// 	for i := 0; i < len(array)-1; i++ {
// 		pointer = (*int)(unsafe.Pointer(memoryAddress))
// 		fmt.Println(*pointer, " ")
// 		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
// 	}

// 	fmt.Println()
// 	pointer = (*int)(unsafe.Pointer(memoryAddress))
// 	fmt.Print("One more : ", *pointer, " ")												//존재 하지않는 배열의 메모리 주소에 접근하여 에러가 발생하지만 따로 에러를 찾아주지 않아 이상한값이 리턴된다.
// 	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
// 	fmt.Println()
// }  