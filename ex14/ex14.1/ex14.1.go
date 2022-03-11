package main

import (
	"fmt"
)

func main() {
	var a int = 500
	var p *int
	//type 앞에 * 붙이는 걸 포인터라고 알수있다 *p와 다르다

	p = &a //&a : 값이 500인 a의 주소값을 p라고 정의
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100 //p(주소)안의 값  = a번지안에 있는 값 즉 500   ,변수는 값이 변하는 메모리 공간   값=메모리 공간
	fmt.Printf("a의 값: %d\n", a)
}
