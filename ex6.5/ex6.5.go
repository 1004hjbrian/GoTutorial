package main

import "fmt"

func main() {
	var x int8 = 127

	fmt.Printf("%d < %d+1: %v\n", x, x, x < +1)
	fmt.Printf("x\t= %4d, ")
}
