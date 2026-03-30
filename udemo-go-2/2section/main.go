package main

import (
	"errors"
	"strings"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil

}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")

	firstName = parts[0]
	lastName = parts[1]

	return

}

func main() {

	result, err := divide(10, 0)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}

	firstName, lastName := splitName("test today")

	println("First Name:", firstName)
	println("Last Name:", lastName)
}
