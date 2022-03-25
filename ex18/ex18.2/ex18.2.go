package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4) //slice에 요소4를 추가한 것, 기존 slice는 바뀌지 않는다

	fmt.Println(slice)
	fmt.Println(slice2)
}
