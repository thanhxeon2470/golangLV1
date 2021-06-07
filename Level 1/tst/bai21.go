package main
import "fmt"

func bai21(n int64) {
	var i int64 = 1
	var S int64 = 0

	for i <= n {
		if n % i == 0 {
			fmt.Print(i," ")
			S += i
		}
		i ++
	}
	fmt.Println("\n",S);
	fmt.Println()

}

func main(){
	var i int64 = 0
	for i < 20{
		bai21(i)
		i += 3
	}
}