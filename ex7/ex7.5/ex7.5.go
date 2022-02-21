package main

import (
	"fmt"
)

func Divide(a, b int) (result int, success bool) { //ex7.4와 다르게 반환타입에 result, success등을 설정했다 -> 괄호output안에 변수처럼 사용할수있다
	if b == 0 {
		result = 0
		success = false
		return //이름 적어줬기에 리턴뒤에 아무것도 적지 않은 것
	}
	result = a / b
	success = true
	return a / b, true //       a/b,true 빼도 돌아간다. 적어도 돌아가는것 확인
}

func main() {
	c, succes := Divide(9, 3)
	fmt.Println(c, succes) //s를 빼도 되는것 확인
	d, success := Divide(9, 0)
	fmt.Println(d, success)

}
