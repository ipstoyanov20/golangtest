package main

import ("fmt"
"github.com/Pallinder/go-randomdata"
)

func presentOptions() {
	fmt.Println("Welcome to the bank")
	fmt.Println(randomdata.PhoneNumber())
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4.Exit")

}