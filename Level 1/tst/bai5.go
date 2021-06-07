package main
import "fmt"

func bai5(n int64) {
	var i int64 = 1
	var S float64 = 0
	if i > n{
		fmt.Println(1,"=",1);
		return
	}

	n = 2 * n;
	fmt.Print(1,"+")
	i += 2
	S = 1
	for i < n{
		fmt.Print(1,"/",i,"+")
		S += 1/float64(i)
		i += 2;
	}
	fmt.Print(1, "/",n + 1);
	S += 1/float64(n + 1);
	fmt.Println("=", S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai5(i)
		i += 3
	}
}