package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("게임을 켠다")
	} else if temp <= 3 {
		fmt.Println("어머니가 화를 내신다")
	} else if temp <= 18 {
		fmt.Println("게임을 끈다!")
	} else {
		fmt.Println("공부를 하자")
	}
}
