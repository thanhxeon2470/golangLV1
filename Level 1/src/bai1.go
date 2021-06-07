package bai1
import "fmt"
func bai1() {
	var n int64
	var i int64 = 1
	var S int64 = 0
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)
	for i < n{
		fmt.Print(i,"+")
		S += i;
		i ++;
	}
	fmt.Print(n);
	S += n;
	fmt.Println("=",S);
}