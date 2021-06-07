package main

import (
	"fmt"
)

func Q8(n int) {
	var sum float32 = 0.0
	var i int = 1

	for ; i < 2*n+1; i++ {
		sum += float32(i) / float32(i+1)
		fmt.Print(i, "/", i+1, "+")
	}

	fmt.Print(i, "/", i+1, "=", float32(2*i-1)/float32(2*i))

	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q8(i)
		i += 3
	}
}
