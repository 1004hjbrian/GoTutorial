package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "hellya mokoko"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	//str1번지를 unsafe패키지안에 있는 pointer타입으로 변환한다. &str1은 *string타입이다
	//이것을 또 reflect패키지안에 있는 StringHeader타입으로 변환 즉, *StringHeader으로 바뀐다
	//*StringHeader는 구조체로 원래 정의된 Len변수와 Data변수로 되어있다
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1) //괄호안 첫째 값은 문자열의 값인 Data, 보관하고 있는 실제 메모리 주소 길이 Len
	fmt.Println(stringHeader2)
}
