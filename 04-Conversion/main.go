package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to Nike store")
	fmt.Println("Rate your AirForce 1 ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating")

	/* The below code would give error, 
	try to execute it and check for the errors,
	 \n is added to the input

	*/

	// numRating, err := strconv.ParseFloat(input, 64)


	// This is the correct code
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating: ", numRating+1)
	}
}
