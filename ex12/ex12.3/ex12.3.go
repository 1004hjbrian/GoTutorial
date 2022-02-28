package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30, 40, 50}
	nums[2] = 300
	for i := 0; i < len(nums); i++ { //len은 내장함수로써 len()괄호 안에 자료구조type에 따라
		//길이를 반환하는 함수 즉, nums는 10부터 50까지 5개있는 배열로 길이 5가 된다
		fmt.Println(nums[i])
	}

	//fmt.Println(nums[i]) for문 밖에서는 선언되지 않았다고 뜸
}
