package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	t := time.Now() //현재시간
	//rand.Seed(100)            //seed함수와 Intn함수(시퀀스)랑 조합해서 전과는 다른 결과값을 도출
	//seed안의 seed가 있어야 매번 다른 값이 나온다 정수넣으면 결국 결과값이 같기때문에 로또형식 x
	rand.Seed(t.UnixNano())
	for i := 0; i < 10; i++ { //10번 뽑는 기능
		fmt.Print(rand.Intn(100), ", ") //rand안에 Intn이라는 함수가 있는데 뒤에 오는 입력값보다
		//0에서 입력값사이의 랜덤값 0부터99까지
	}

}
