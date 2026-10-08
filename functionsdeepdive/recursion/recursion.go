package recursion


import "fmt"

func main() {
	fact := factorial(5)

	fmt.Print(fact)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	// result := 1
	// for i:= 1; i <= n; i++{
	// 	result = result * i
	// }
	// return result
	return n * factorial(n-1)
}
