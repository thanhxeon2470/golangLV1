package lv1

import (
	"fmt"
)

func Q2() {
	var n int
	var sum int
	var i int = 1

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		sum += i
		fmt.Print(i, "^2+")
	}

	fmt.Print(n, "^2=", sum+n)

}
