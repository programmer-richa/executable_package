// Package main specifies that this is an executable package.
package main

import (
	"fmt"

	"github.com/programmer-richa/utility_package/constants"
	"github.com/programmer-richa/utility_package/constants/description"
)

// It is the entry point of the program
func main() {
	fmt.Println("An executable package is the one that has one package main and a main func defined in it.")
	fmt.Printf("Author: %s\n", description.Author)
	fmt.Printf("Author: %s\n", description.Date)
	fmt.Printf("Description: %s\n", description.Description)
	fmt.Printf("Login success message: %s\n", constants.LoginSucc)
}
