package lv1

import (
	"fmt"
)

func Q25() int {
	var n int
	var sum int = 0
	var i int = 2

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)
	if n == 0 {
		print("0 ")
		return 0
	}
	for ; i <= n/2; i += 2 {
		if n%i == 0 {
			sum += i
		}
	}
	fmt.Print(sum)
	return sum
}
