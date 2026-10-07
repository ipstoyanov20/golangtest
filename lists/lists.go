package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)

	prices = prices[1:]

	fmt.Println(prices)

}

// func main() {

// 	var productNames [4]string = [4]string{"A Book", }
// 	prices := [4]float64{1.2, 2.3, 3.4, 9.99}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featurePrices := prices[1:]

// 	featurePrices[0] = 100.0

// 	highlightedPrices := featurePrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
