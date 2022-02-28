package main

import (
	"fmt"
)

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, //여기도 쉼표
	}

	for _, arr := range a { //선언된 이중배열 a에 대해 번호는 생략 값만 호출 ,여기서 arr 타입은 [5]int , a[0],a[1]이렇게 첫번째 묶음 두번째 묶음
		for _, v := range arr { //위의 값 arr에 대해 번호 생략 값만 호출 , 여기서 v타입은 int이다. 묶음 안의 원소들을 차례로 호출
			fmt.Print(v, " ") //v값 호출하고 한칸 띄고
		}
		fmt.Println() //문장 한칸 띈다
	}

}

//쉽게 얘기하면 range a 는 행 range arr 는 열이 되겟다
