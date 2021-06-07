package main
import "fmt"

func bai9(n int64) {
	var i int64 = 1
	var S int64 = 1
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

func main(){
	var i int64 = 0
	for i < 20{
		bai9(i)
		i += 3
	}
}