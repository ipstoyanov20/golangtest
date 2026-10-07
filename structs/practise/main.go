package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs/practise/note"
	"example.com/structs/practise/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	result := add(2, 3)

	fmt.Print(result)

	printSomething(2)
	printSomething(2.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	printSomething(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)
}

func add[T int | float64 | string](a, b T) T {

	return a + b
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if !ok {
		fmt.Printf("%v %v\n", intVal, ok)
	}
	floatVal, ok := value.(float64)

	if !ok {
		fmt.Printf("%v %v\n", floatVal, ok)
	}

	stringVal, ok := value.(string)

	if !ok {
		fmt.Printf("%v %v\n", stringVal, ok)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("This is an integer", value)
	// case float64:
	// 	fmt.Println("This is a float", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save the note!")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
