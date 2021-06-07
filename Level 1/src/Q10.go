package lv1

import (
	"fmt"
	"math"
)

func Q10() {
	var n int
	var x float64

	fmt.Print("Nhap vao x: ")
	fmt.Scan(&x)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	fmt.Print(x, "^", n, "= ", math.Pow(x, float64(n)))
	// for ; i < n; i++ {
	// 	sum += i
	// 	fmt.Print(i, "+")
	// }

	// fmt.Print(n, "=", sum+n)

}
