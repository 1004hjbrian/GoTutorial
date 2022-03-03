package main

import (
	"fmt"
	"unsafe" //표준패키지로써 안전하지 않는 함수들은 제거
)

type User struct {
	Age   int32   //4byte
	Score float64 //8byte
} //12byte

func main() {
	var a int = 8
	user := User{23, 75.5}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a)) //user의 메모리 공간의 사이즈를 반환
}

//첫 값이 12가 아니고 16바이트 나온이유는 메모리 정렬할때 Age 와 Score 변수 사이에 빈 공간이 생기기에 실제로 다름
