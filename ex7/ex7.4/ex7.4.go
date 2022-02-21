package main

import (
	"fmt"
)

func Divide(a, b int) (int, bool) { //int a 는 int로 반환 , int b는 bool타입으로 반환하는 Divide함수 (x) => int ,bool은 return 반환타입인듯
	if b == 0 {
		return 0, false
	} //b가 0이면 0을 반환하면 false

	return a / b, true //b가 0이 아닐때 a나누기b를 하고 true를 반환
}

func main() {
	c, success := Divide(9, 3) // c와 success를 정의, 매개변수 a=9 b=3을 넣은 Divide 함수를 실행
	fmt.Println(c, success)    //반환값 출력 (c는 함수안의 첫번째 return값(int), success는 두번재 bool값 ) **c,success 변수들은 ouput에 맞는 값
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
