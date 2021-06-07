package lv1

import (
	"fmt"
)

func Q8() {
	var n int
	var sum float32 = 0.0
	var i int = 1

	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < 2*n+1; i++ {
		sum += float32(i) / float32(i+1)
		fmt.Print(i, "/", i+1, "+")
	}

	fmt.Print(i, "/", i+1, "=", float32(2*i-1)/float32(2*i))

}
