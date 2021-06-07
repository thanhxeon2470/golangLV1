package main
import "fmt"
// import "math"

func main() {
	var n int64
	var i int64 = 1
	var S float64 = 0
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)
	if i > n{
		fmt.Println(n," = ",n);
		return
	} 
	for i < n{
		fmt.Print(i,"/",i+1,"+")
		S += float64(i)/float64(i+1)
		i+=1;
	}
	fmt.Print(i,"/",n+1)
	S += float64(i)/float64(n+1)
	fmt.Println("=",S);
}