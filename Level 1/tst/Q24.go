package main

import (
	"fmt"
)

func Q24(n int) {
	var i int = 1

	if n == 0 {
		println("0 khong co uoc so !!")
		return
	}
	for ; i <= n/2; i += 2 {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	if n%2 != 0 {
		fmt.Print(n)
	}

	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q24(i)
		i += 3
	}
}
