package main
import "fmt"
// import "math"

func main() {
	var n int64
	var i int64 = 1
	var S float64 = 0
	var j int64
	fmt.Print("Nhap n vao day: ")
	fmt.Scan(&n)

	for i < n{
		var tmp float64 = 1
		j = 1
		for j < i{
			fmt.Print(j,"+")
			tmp += j
			j++
		}
		fmt.Print(i,"+")
		tmp += i
		S += 1 / tmp
		i += 1;
	}

	j = 1
	var tmp int64 = 1
	for j < n{
		fmt.Print(j,".")
		tmp*=j
		j++
	}
	fmt.Print(n)
	tmp*=n
	S += tmp
	
	fmt.Println("=",S);
}