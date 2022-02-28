package main

import (
	"fmt"
)

func main() {
	var t [5]float64 = [5]float64{23.0, 25.5, 27.7, 30.2, 34.5}

	for i, v := range t { //t를 range키워드로 for문 안에서 쓰이고 각 요소를 순회하면서 값 두개를 반환한다. 배열의
		//경우에는 첫번째값은 index , 두번째는 요소의 값을 반환
		fmt.Println(i, v)
	}

}
