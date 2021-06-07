package main

import (
	"fmt"
	"math"
)

func CalcFac(x int) float64 {
	var i int
	var sum int = 1
	if x <= 1 {
		return 1
	}
	for i = 1; i < x; i++ {
		sum *= i
	}

	return float64(sum)
}
func Q18(x float64, n int) {
	var sum float64 = 0.0
	var i int = 0

	for ; i < 2*n; i += 2 {
		fmt.Print(x, "^", i, "/", i, "!")

		sum += math.Pow(x, float64(i)) / CalcFac(i/2)
		fmt.Print("+")
	}

	fmt.Print(x, "^", i, "/", i, "!")
	sum += math.Pow(x, float64(i)) / CalcFac(i/2)
	fmt.Print(" = ", sum)

	fmt.Println()
}

func main(){
	var i int = 0
	for i < 20{
		Q18(2, i)
		i += 3
	}
}
