package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("HJP", 23)

	fmt.Println(userPointer)
}

//p1 := &Data{}  여기서 p1의 타입은 *Data 포인터 Data타입이다
//생성자로 똑같은 작업을 간단하게 표현하자면
//var p2 = new(Data)  ==>우항은 위의 우항과 같은 개념이다
//대신에 new는 다른 값으로 넣을 수 없어서 다른 값으로 초기화가 불가능하다는 단점이 있다.

//초기화란 변수의 초깃값 할당을 의미
