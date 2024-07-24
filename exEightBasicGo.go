package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	capitalizedText := strings.ToUpper(text)
	tinyText := strings.ToLower(text)

	fmt.Println("Original Text:", text)
	fmt.Println("In capital letters:", capitalizedText)
	fmt.Println("In lowercase:", tinyText)
}
