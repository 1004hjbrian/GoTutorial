//go언어의 일관성 때문에 main안의 slice와 addNum의 slice는 서로 다른
//인스터스값 -> 해결법 두가지 포인터 or return slice
package main

import "fmt"

//slice는 동적배열
func addNum(slice []int) {
	slice = append(slice, 4)
}

func main() {
	fmt.Println()
}
