package main

import (
	"fmt"
)

func Q23(n int) int {
	var count int
	var i int = 1

	if n == 0 {
		println("0 khong co uoc so !!")
		return 0
	}
	for ; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	fmt.Print(count)
	fmt.Println()
	return count
}

func main(){
	var i int = 0
	for i < 20{
		Q23(i)
		i += 3
	}
}
