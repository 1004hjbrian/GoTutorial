//재귀함수
package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1) //함수로 다시 돌아가므로 if 문으로 돌아가는 것 //함수 안의 함수 형태
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}

//printNo(3) 으로 시작하여 fmt.Println(n)으로  '3' 이 먼저 출력
//printNo(n-1)로 printNo(2)로 '2' , printNo(1)로 '1', 10번째줄에서 printNo(0)함수를 다시 시작하여
//if n==0 으로 return 여기서 return은 돌려주는 값 없이 printNo(n-1) 돌아가는 것 같다.

// printNo(3){
// 	3
//    printNo(2){
// 		2
// 		  printNo(1){
// 			1
// 			printNo(0){
// 				after1(x)      //return 함수종료
// 			}
//             after2(x)     //after1
// 		  }
//           after3(x)     //after2
// 	   }
//                    //after3 일 확률이 더 높을듯
// 	}
