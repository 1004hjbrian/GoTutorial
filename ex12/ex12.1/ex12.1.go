package main

import (
	"fmt"
)

func main() {
	var s [5]float64 = [5]float64{24, 23.3, 25.5, 26.7, 34.1}
	//	또는 s := [5]float64{24, 23.3, ....} 선언가능
	fmt.Println(s[2])
}
