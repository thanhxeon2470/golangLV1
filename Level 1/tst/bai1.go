package main
import "fmt"
func bai1(n int64) {

	var i int64 = 1
	var S int64 = 0

	for i < n{
		fmt.Print(i,"+")
		S += i;
		i ++;
	}
	fmt.Print(n);
	S += n;
	fmt.Println("=",S);
}

func main(){
	var i int64 = 0
	for i < 20{
		bai1(i)
		i += 3
	}
}