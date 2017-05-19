package main

import "fmt"

func main() {

	num := 1

	if num > 3 {
		fmt.Print("manay")
	}

	// Go is picky, "else" must be on the same line as closing if bracket
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
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
