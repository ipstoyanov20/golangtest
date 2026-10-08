package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4}
	sum := sumup(1, 10, 15, 40, -5)
	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)

}

func sumup(start int, numbers ...int) int {

	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum

}
