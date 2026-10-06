package main

import (
	"example.com/bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println(" error123", err, "-----")
		panic("Can't continue, sorry")
	}
	presentOptions()
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your balance:", accountBalance)
	case 2:
		fmt.Println("Your BALANCE 2:", accountBalance)
		fileops.WriteFloatToFile(accountBalance+200, accountBalanceFile)
	case 3:
		fmt.Println("Your BALANCE 3:", accountBalance)
	case 4:
		fmt.Println("Your BALANCE 4:", accountBalance)
	case 5:
		fmt.Println("Your BALANCE 5:", accountBalance)
	}

	fmt.Println("Your choice:", choice)
}
