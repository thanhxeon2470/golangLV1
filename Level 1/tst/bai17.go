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

func bai17(x float64,n int64) {
	var i int64 = 2
	var S float64

	if n < 1{
		fmt.Println("n phai lon hon 0!");
		return
	}
	if n == 1{
		fmt.Println(x,"=",x);
		return
	}
	fmt.Print(x,"+")
	S = x
	for i < n{
		fmt.Print(x,"^", i, "/", i, "!+")
		S += math.Pow(x,float64(i)) / float64(factorial(i))
		i+=1;
	}
	fmt.Print(x,"^", n, "/", n, "!")
	S += math.Pow(x,float64(n)) / float64(factorial(n));
	fmt.Println(" =",S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai17(2,i)
		i += 3
	}
}