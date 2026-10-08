package main

import "fmt"

func printHobbiesArrays() *[]string {
	hobbies := []string{"reading", "swimming", "cooking"}
	//print it using for loop auto
	for _, hobby := range hobbies {
		fmt.Println(hobby)
	}
	return &hobbies
}

func printSlices2(h *[]string) *[]string {

	fmt.Println("First element:", (*h)[0])
	fmt.Println("Second and third element combined:", (*h)[1:])
	return h

}

func printSlices3(h *[]string) *[]string {
	fmt.Println("the first and second elements:", (*h)[:2])
	return h
}

func printSlices4(h *[]string) *[]string {
	fmt.Println("Re-slice the slice from (3) and change it to contain the second and last element of the original array.:", (*h)[1:3])
	return h
}

type Product struct{
	id string
	title string
	price float64
}

func main() {
	slice := printSlices2(printHobbiesArrays())

	slice = printSlices3(slice)

	printSlices4(slice)

	goals := []string{"learn Go", "Make profit for the CBT projects"}
	fmt.Println(goals)


	goals[1] = "I don't want to change it"
	goals = append(goals, "Be a better programmer")

	fmt.Println(goals)

	products := []Product{
		{"jfiwjrf83", "snowboard", 12.32},
		{"kjfuhf737", "book", 12.32},
	}
	
	fmt.Println(products)
	
	newProduct := Product{
		"kjfuhf7372312", "book2", 12.32,
	}
	
	products = append(products, newProduct)
	fmt.Println(products)
	

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
