package main

import "fmt"

type str string

func (text str) log() {
	fmt.Print(text)
}

func main() {
 var name str = "Ivo"
 name.log()
}
