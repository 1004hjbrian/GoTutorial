package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1
	B int  //8
	C int8
	D int
	E int8
} //19바이트  //실제는 40바이트 패딩때문에

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
