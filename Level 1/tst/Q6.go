package main

import (
	"fmt"
)

func Q6(n int) {
	var sum float32 = 0.0
	var i int = 1

	fmt.Println(1.0 / 2)

	for ; i < n; i++ {
		sum += 1.0 / float32(((i + 1) * i))
		fmt.Print("1/", i, "x", i+1, "+")
	}

	fmt.Printf("1/%dx%d = %f", i, i+1, sum+1.0/float32(((i+1)*i)))
	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q6(i)
		i += 3
	}
}
