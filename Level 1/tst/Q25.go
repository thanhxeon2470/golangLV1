package main

import (
	"fmt"
)

func Q25(n int) int {
	var sum int = 0
	var i int = 2

	if n == 0 {
		print("0 ")
		return 0
	}
	for ; i <= n/2; i += 2 {
		if n%i == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
	return sum
}

func main(){
	var i int = 0
	for i < 20{
		Q25(i)
		i += 3
	}
}

