package main
import "fmt"

func bai11(n int64) {
	var i int64 = 1
	var S int64 = 0
	var j int64

	for i < n{
		var tmp int64 = 1
		j = 1
		for j < i{
			fmt.Print(j,".")
			tmp*=j
			j++
		}
		fmt.Print(i,"+")
		tmp*=i
		S += tmp
		i+=1;
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

func main(){
	var i int64 = 0
	for i < 20{
		bai11(i)
		i += 3
	}
}