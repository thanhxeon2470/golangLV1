package main
import "fmt"

func bai3(n int64) {
	var i int64 = 1
	var S float64 = 0

	for i < n{
		fmt.Print(1,"/",i,"+")
		S += 1/float64(i)
		i ++;
	}
	fmt.Print(1,"/",n);
	S += 1/float64(n);
	fmt.Println("=",S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai3(i)
		i += 3
	}
}