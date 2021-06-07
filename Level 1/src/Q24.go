package lv1

import (
	"fmt"
)

func Q24() {
	var n int

	var i int = 1

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)
	if n == 0 {
		print("0 khong co uoc so !!")
		return
	}
	for ; i <= n/2; i += 2 {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	if n%2 != 0 {
		fmt.Print(n)
	}

}
