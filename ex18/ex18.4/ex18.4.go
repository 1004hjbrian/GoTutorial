package main

import (
	"fmt"
)

func changeArray(array2 [5]int) {
	array2[2] = 200 //array2의 3번째값에 해당되는값 200

	//매개변수 array2 변수만을 써야한다 아니면 빨간줄 생김
	//go는 일관성, 함수를 호출할때 매개변수는 rvalue로 값으로 인식한다 즉 대입연산자 마냥 사용한것과 같다
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	//func change 의 array 와 func main 의 array 이름이 같아도
	//서로 다른 인스턴스이다.

	fmt.Println(array)
	fmt.Println(slice)
}
