package main

import (
	"fmt"
)

func main() {
	a := 4

	switch a {
	case 1:
		fmt.Println("a==1")
		break //go언어에선 안써도 알아서 break, break안하려면 fallthrough를 쓰면 된다
	case 2:
		fmt.Println("a==2")
	case 3:
		fmt.Println("a==3")
		fallthrough
	case 4:
		fmt.Println("a==4")
	default:
		fmt.Println("a !=1,2,3")

	}
	fmt.Println()
}

//====>>>>>>>값이 3이아니어도 fallthrough가 있으면 다음값도 출력이 된다고 기억하자.
