package lv1

import (
	"fmt"
	"math"
)

func Q14() {
	var n int
	var sum float64 = 0.0
	var i int = 1
	var x float64

	fmt.Print("Nhap vao x: ")
	fmt.Scan(&x)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < 2*n+1; i += 2 {
		sum += math.Pow(x, float64(i))
		fmt.Print(x, "^", i, "+")
	}

	fmt.Print(x, "^", i, "=", sum+math.Pow(x, float64(i)))

}
