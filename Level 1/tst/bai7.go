package main
import "fmt"

func bai7(n int64) {
	var i int64 = 1
	var S float64 = 0
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


func main(){
	var i int64 = 0
	for i < 20{
		bai7(i)
		i += 3
	}
}