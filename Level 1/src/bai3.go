package main
import "fmt"
// import "math"

func main() {
	var n int64
	var i int64 = 1
	var S float64 = 0
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)

	for i < n{
		fmt.Print(1,"/",i,"+")
		S += 1/float64(i)
		i ++;
	}
	fmt.Print(1,"/",n);
	S += 1/float64(n);
	fmt.Println("=",S);
}