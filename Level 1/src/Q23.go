package lv1

import (
	"fmt"
)

func Q23() int {
	var n int
	var count int
	var i int = 1

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)
	if n == 0 {
		print("0 khong co uoc so !!")
		return 0
	}
	for ; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	fmt.Print(count)
	return count

}
