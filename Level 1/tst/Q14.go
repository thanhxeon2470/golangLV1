package main

import (
	"fmt"
	"math"
)

func Q14(x float64, n int) {
	var sum float64 = 0.0
	var i int = 1


	for ; i < 2*n+1; i += 2 {
		sum += math.Pow(x, float64(i))
		fmt.Print(x, "^", i, "+")
	}

	fmt.Print(x, "^", i, "=", sum+math.Pow(x, float64(i)))

	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q14(2, i)
		i += 3
	}
}
