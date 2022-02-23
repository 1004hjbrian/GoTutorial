package main

import "fmt"

type ColorType int

const (
	Red  ColorType = iota
	Blue ColorType = iota //그냥 Blue라고 적어도 된다
	Green
	Yellow
)

func ColorToString(color ColorType) string {
	switch color {
	case Red:
		return "RED"
	case Blue:
		return "BLUE"
	case Green:
		return "GREEN"
	case Yellow:
		return "YELLOW"
	default:
		return "Undefined"
	}
}

func GetMyFavoriteColor() ColorType {
	return Blue //output이 ColorType이기에 Blue를 쓸 수 있는것이다.
}

func main() {
	fmt.Println("내가 좋아하는 색은", ColorToString(GetMyFavoriteColor())) //color안의 () 빼고 돌려보자->함수이므로 안됨
	//앞의 colortostring 빼서 돌리면 1이 나온다. iota 01234 중 두번째로 선언된 blue는 1이다
	//getmyfavorite 안에 colortostring을 넣어보자 ===>>>> argument가 없다고 콜이 안된다 한다
	//===>'Blue' argument가 있는 선언된 ColorToString 함수안에서 getmyFavoriteColor를 불러야 된다.
	//다른 방법으로는 colortostring의 매개변수란에 1을 쓰면 Blue가 나온다
}
