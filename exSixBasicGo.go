package main

import "fmt"

func main() {
	names := []string{"Jacob", "Mary", "Peter"}

	for _, name := range names {
		fmt.Println("Hello,", name)
	}
}
