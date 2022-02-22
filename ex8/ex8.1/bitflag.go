package main

import (
	"fmt"
)

type Room uint8

const (
	MasterRoom uint8 = 1 << iota
	BathRoom
	Kitchen
	LivingRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
	if IsLightOn(rooms, Kitchen) {
		fmt.Println("부엌에 불을 켠다")
	}

}

func main() {
	var rooms uint8 = 0                 //처음엔 모든 비트가 0으로 설정
	rooms = SetLight(rooms, MasterRoom) //masteroom 0000 0001로 바꾸어라
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, LivingRoom)

	rooms = ResetLight(rooms, BathRoom) //reset이므로 0으로 바꾼다는것

	TurnLights(rooms)
	fmt.Println()
}
