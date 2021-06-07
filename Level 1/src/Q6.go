package lv1

import (
	"fmt"
)

func Q6() {
	var n int
	var sum float32 = 0.0
	var i int = 1

	fmt.Println(1.0 / 2)
	fmt.Print("Nhap vao n: ")
	fmt.Scan(&n)

	for ; i < n; i++ {
		sum += 1.0 / float32(((i + 1) * i))
		fmt.Print("1/", i, "x", i+1, "+")
	}

	fmt.Printf("1/%dx%d = %f", i, i+1, sum+1.0/float32(((i+1)*i)))

}
