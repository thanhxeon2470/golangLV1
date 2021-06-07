package lv1

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
func Q18() {
	var n int
	var sum float64 = 0.0
	var i int = 0
	var x float64

	fmt.Print("Nhap vao x: ")
	fmt.Scan(&x)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < 2*n; i += 2 {
		fmt.Print(x, "^", i, "/", i, "!")

		sum += math.Pow(x, float64(i)) / CalcFac(i/2)
		fmt.Print("+")
	}

	fmt.Print(x, "^", i, "/", i, "!")
	sum += math.Pow(x, float64(i)) / CalcFac(i/2)
	fmt.Print("=", sum)

}
