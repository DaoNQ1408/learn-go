package main

import "fmt"

func myPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something bad happened in myPanic")
	}

	fmt.Println("this function will not panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	myPanic(true)
}

func main() {
	recoverable()
}
