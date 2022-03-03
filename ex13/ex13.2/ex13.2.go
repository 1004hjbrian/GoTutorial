package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User //위의 User구조체
	VIPLevel int
	Price    int
}

func main() {
	user := User{"HyeonJun", "hj", 23} //user라는 변수로 유저 목록을 정의함
	vip := VIPUser{
		User{"michael", "mj", 50}, //User 타입 초기화   여기가 핵심
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이:%d\n,   VIP레벨: %d 가격: %d\n",
		vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)

}
