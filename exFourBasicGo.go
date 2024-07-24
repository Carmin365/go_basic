package main

import "fmt"

func main() {
	fruit := "apple"

	switch fruit {
	case "banana":
		fmt.Println("You chose banana!")
	case "apple":
		fmt.Println("You chose apple!")
	case "orange":
		fmt.Println("You chose orange!")
	default:
		fmt.Println("Fruit not found!")
	}
}
