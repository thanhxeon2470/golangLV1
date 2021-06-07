package lv1

import (
	"fmt"
	"math"
)

func Q12() {
	var n int
	var sum float64 = 0.0
	var i int = 1
	var x float64

	fmt.Print("Nhap vao x: ")
	fmt.Scan(&x)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		sum += math.Pow(x, float64(i))
		fmt.Print(x, "^", i, "+")
	}

	fmt.Print(x, "^", i, "=", sum+math.Pow(x, float64(i)))

}
