package lv1

import (
	"fmt"
)

func Q4() {
	var n int
	var sum float32 = 0.0
	var i int = 1

	fmt.Println(1.0 / 2)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		sum += 1.0 / float32((2 * i))
		fmt.Print("1/", i*2, "+")
	}

	fmt.Printf("1/%d = %f", 2*n, sum+1.0/float32(i*2))

}
