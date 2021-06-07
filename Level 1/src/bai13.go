package main
import ("fmt"
		"math")
// import "math"

func main() {
	var x float64
	var n int64
	var i int64 = 1
	var S int64 = 0
	fmt.Print("Nhap x vao day: ")
	fmt.Scan(&x)
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)

	for i < n{
		fmt.Print(x,"^",2*i,"+")
		S += int64(math.Pow(x, float64(2*i)))
		i++;
	}
	fmt.Print(x,"^",n*2)
	S += int64(math.Pow(x, float64(n*2)))
	fmt.Println("=",S);
}