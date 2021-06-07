package main
import "fmt"
// import "math"

func main() {
	var n int64
	var i int64 = 1
	var S int64 = 1
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)
	if i >= n{
		fmt.Println(n,"=",n);
		return
	} 
	for i < n{
		fmt.Print(i,"x")
		S *= i
		i+=1;
	}
	fmt.Print(n)
	S *= n
	fmt.Println(" =",S);
}