package main

import (
	"fmt"
	"math"
)

func Q10(x float64, n int) {
	
	fmt.Print(x, "^", n, "= ", math.Pow(x, float64(n)))
	// for ; i < n; i++ {
	// 	sum += i
	// 	fmt.Print(i, "+")
	// }

	// fmt.Print(n, "=", sum+n)
	fmt.Println()

}

func main(){
	var i int = 0
	for i < 20{
		Q10(2, i)
		i += 3
	}
}

