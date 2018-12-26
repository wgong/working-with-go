/**
 * control-structures.go
 * If, Else, Switch and Conditionals
 */

// standard stuff
package main

import "fmt"

func main() {

	// define an int variable
	num := 0

	fmt.Println("num=", num)

	// Go is not picky, conditionals don't require parentheses
	if num > 3 {
		fmt.Println("Many")
	}

	// Go is picky, "else" must be on the same line as closing if bracket
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Neither 1 nor 2")
	}

	// Switch statement takes conditionals and auto breaks
	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}

}
