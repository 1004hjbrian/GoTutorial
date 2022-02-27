package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("숫자입력해라:")
		var number int
		_, err := fmt.Scanln(&number) //scan은 2개의 값을 읽어오는데 갯수와 error다
		//하지만 갯수에 변수를 할당안했으니 n이 아닌 빈칸을 의미하는 '_'을 넣었다
		if err != nil { //에러가 아닌게 아니라면 즉 에러가 있다면
			fmt.Println("숫자로 입력해라잉")

			//키보드 버퍼를 지웁니다.
			stdin.ReadString('\n') //\n는 문자 하나를 의미
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break // 나머지가2일때 혹은 짝수면 break
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
