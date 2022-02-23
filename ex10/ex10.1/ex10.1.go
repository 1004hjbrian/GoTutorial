//if와 다르게 switch는 조건이 아닌  값을 비교해
//지정한 case 값에 따라 결과를 출력하는데 목적을 둔다
package main

import (
	"fmt"
)

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a > 3")
	}
	fmt.Println()
}
