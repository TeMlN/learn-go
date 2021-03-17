package main //cbr is "call by value" 

import (
	"fmt"
)

func pointInc(i *int) {
	*i = *i + 1
}

func inc(i int) {
	i = i + 1
}

func main() {
	i := 10
	inc(i)
	fmt.Println(i)		//실행 결과에서 봤듯이 함수 내부에서는 값을 변경할 수 없다. 그러므로 &연산자를 이용하여야 한다.

	i = 10
	pointInc(&i)
	fmt.Println(i)		//첫번째 결과에 달리 반면 연산이 잘된것을 볼수가 있다. 함수 내부에서 값을 변경하려면 포인터를 쓰도록하자.


}					

