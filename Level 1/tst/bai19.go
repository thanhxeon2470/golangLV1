package main
import "fmt"
import "math"


func factorial(x int64) int64 {
	var n int64 = 1
	if x == 0{
		return n
	}
	var i int64 = 2
	for i <= x{
		n *= i
		i++
	}
	return n
}

func bai19(x float64, n int64) {
	var i int64 = 3
	var S float64

	if n < 0{
		fmt.Println("n phai lon hon hoac bang 0");
		return
	}
	if n == 0{
		fmt.Println("1+",x,"=",1+x)
		return
	}
	fmt.Print("1+",x,"+")
	S = x + 1
	n=2*n
	for i < n {
		fmt.Print(x,"^", i, "/", i, "!+")
		S += math.Pow(x,float64(i)) / float64(factorial(i))
		i += 2;
	}
	fmt.Print(x,"^", n+1, "/", n+1, "!")
	S += math.Pow(x,float64(n+1)) / float64(factorial(n+1));
	fmt.Println(" =",S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai19(2,i)
		i += 3
	}
}