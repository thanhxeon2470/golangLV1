package main

import (
	"fmt"
	"math"
)

func PrintSum(x int) float64 {
	var i int
	var sum int = 0
	for i = 1; i < x; i++ {
		sum += i
		fmt.Print(i, "+")
	}
	fmt.Print(x)
	if x == 1 {
		sum = 1
	}
	return float64(sum)
}
func Q16(x float64, n int) {
	var sum float64 = 0.0
	var i int = 1

	for ; i < n; i++ {
		fmt.Print(x, "^", i, "/")
		var tmp float64 = PrintSum(i)
		sum += math.Pow(x, float64(i)) / tmp
		fmt.Print("+", tmp)
	}

	fmt.Print(x, "^", i, "/")
	sum += math.Pow(x, float64(i)) / PrintSum(i)
	fmt.Print("=", sum)
	fmt.Println()

}

func main(){
	var i int = 0
	for i < 20{
		Q16(2, i)
		i += 3
	}
}
