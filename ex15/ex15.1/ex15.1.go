package main

import "fmt"

func main() {
	poet1 := "나는 HJ이고 덩크를 목표로 하고 있다 개발자로서도\n성공하고 싶다"

	poet2 := `나는 HJ이고 덩크를 목표로 하고 있다
	 개발자로서도 성공하고 싶다`
	fmt.Println(poet1, poet2)
}
