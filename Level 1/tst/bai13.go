package main
import ("fmt"
		"math")
// import "math"

func bai13(x float64, n int64) {
	var i int64 = 1
	var S int64 = 0

	for i < n{
		fmt.Print(x,"^",2*i,"+")
		S += int64(math.Pow(x, float64(2*i)))
		i++;
	}
	fmt.Print(x,"^",n*2)
	S += int64(math.Pow(x, float64(n*2)))
	fmt.Println("=",S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai13(2,i)
		i += 3
	}
}