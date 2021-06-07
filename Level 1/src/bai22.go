package main
import "fmt"

func main() {
	var n int64
	var i int64 = 1
	var S int64 = 1

	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)

	for i <= n {
		if n % i == 0 {
			fmt.Print(i," ")
			S *= i
		}
		i ++
	}
	fmt.Println("\n",S);
}