// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func printStats(mem runtime.MemStats) {
// 	runtime.ReadMemStats(&mem)
// 	fmt.Println("mem.Alloc:", mem.Alloc)
// 	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
// 	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
// 	fmt.Println("mem.NumGC:", mem.NumGC)
// 	fmt.Println("-----")
// }

// func main() {
// 	var mem runtime.MemStats
// 	printStats(mem)

// 	for i := 0; i < 10; i++ {                   //거대한 Go슬라이스를 생성.
// 		s := make([]byte, 50000000) 			//방대한 메모리를 할당해서 가비지 컬렉터를 호출.
// 		if s == nil {							//만약 s가 초가화 되지 않았다면.
// 			fmt.Println("Operation failed!")	//Operation failed를 출력해라.
// 		}
// 	}
// 	printStats(mem)

// 	for i:=0; i<10; i++ {
// 		s := make([]byte, 100000000)
// 		if s == nil {
// 			fmt.Println("Operation failed!")
// 		}
// 		time.Sleep(5 * time.Second)
// 	}
// 	printStats(mem)
// } 