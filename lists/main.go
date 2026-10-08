package main

import "fmt"

type floatMap map[string]float64


func (m floatMap) output(){
	fmt.Print(m)
}

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()


	fmt.Println("Usernames")
	for index, val := range userNames{
		fmt.Println("Index:", index)
		fmt.Println("Value:", val)
		
		
	}
	
	fmt.Println("courseRatings")
	for i, v := range courseRatings{
		fmt.Println("Key", i)
		fmt.Println("Value", v)
	}

}
