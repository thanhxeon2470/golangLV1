package main
import "fmt"

func bai20(n int64) {
	var i int64 = 1

	for i <= n {
		if n % i == 0 {
			fmt.Print(i," ")
		}
		i ++
	}
	fmt.Println()
}

func main(){
	var i int64 = 0
	for i < 20{
		bai20(i)
		i += 3
	}
}