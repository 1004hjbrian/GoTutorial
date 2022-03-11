package main

import "fmt"

type Data struct {
	value int      //8
	data  [200]int //1608바이트
} //1608바이트  다 초기값 0으로 가지고 있다

func ChangeData(arg *Data) { //파라미터 arg  구조체Data 타입의 주소값으로 arg를 인자로
	//구조체의 양식에 맞춰 arg의 vaue 값과 data 값을 정의
	arg.value = 999     //원래는 (*arg)이 원칙인데 go에선 편하게 arg해도 되게 설계함
	arg.data[100] = 999 //101번째? 99번째라고 생각한다
}

func main() {
	var data Data

	ChangeData(&data) //func ChangeData 호출해서 인자로 data의 주소값을 호출
	// -> ChangeData(arg Data)라 Data 타입의 인자만 부를 수 있다

	//나의 생각: 변수 data가 arg에 들어간거

	//현재 var data가 ChangeData 들어가있는 상태

	fmt.Printf("value = %d\n", data.value)         // var data의 value 값  arg에 data가 들어갔으니 data.value
	fmt.Printf("data[100] = %d\n", data.data[100]) //
}

//함수안에 들어간것은 우항이다 (좌항 주항)  ==> fmt.Println(여기 말하는거)

//func ChangeData(arg Data) 이고 func main(){ChangeData(data)일때 }
//값이 변하지 않고 초깃값 0 들이 나오는 이유 ==> 대입한 var data 의 값과 ChangeData
//의 arg값은 관련이 없다고 한다  data는 바뀐적이 없기 때문이다
