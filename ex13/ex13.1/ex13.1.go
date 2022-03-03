package main

import (
	"fmt"
)

type House struct { //house라는 이름의 구조체 타입
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House //var house-> 변수인데 House구조체의 house라는 변수를 정의

	//house에 대한 정보 (구조체의 양식에 따른 정보 입력)
	house.Address = "서울시 강남구"
	house.Size = 10
	house.Price = 80
	house.Category = "오피스텔"

	fmt.Println(house) //혹은 fmt.Printf("%v\n",house)  ==>구조체에선 %d 대신 %v를 쓴다
	//또다른 표현 fmt.Printf("주소:%s 사이즈:%d평 가격:%d억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
	fmt.Println("guswns 집 구하기 작전")
}
