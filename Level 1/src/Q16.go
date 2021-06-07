package lv1

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
func Q16() {
	var n int
	var sum float64 = 0.0
	var i int = 1
	var x float64

	fmt.Print("Nhap vao x: ")
	fmt.Scan(&x)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		fmt.Print(x, "^", i, "/")
		var tmp float64 = PrintSum(i)
		sum += math.Pow(x, float64(i)) / tmp
		fmt.Print("+", tmp)
	}

	fmt.Print(x, "^", i, "/")
	sum += math.Pow(x, float64(i)) / PrintSum(i)
	fmt.Print("=", sum)

}
