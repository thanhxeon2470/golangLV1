package main

import (
	"fmt"
)

func Q2(n int) {
	var sum int
	var i int = 1

	for ; i < n; i++ {
		sum += i
		fmt.Print(i, "^2+")
	}

	fmt.Print(n, "^2=", sum+n)
	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q2(i)
		i += 3
	}
}
