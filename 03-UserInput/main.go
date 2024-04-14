package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome"
	fmt.Printf(welcome)

	// taking the input from the user: https://pgk.go.dev

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a rating for a product: ")

	// comma okay || err okay syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of rating %T, ", input)

}
