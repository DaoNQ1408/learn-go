package main

import "fmt"

func simpleDefer() {
	fmt.Println("Start of function")
	defer fmt.Println("End of function") // This will be executed at the end of the function, just before it returns
	fmt.Println("Middle of function")
}

func lifoSimpleDefer() {
	fmt.Println("Start of function")
	defer fmt.Println("first defer") // multi defer will be executed in LIFO order
	defer fmt.Println("second defer")
	fmt.Println("Middle of function")
}

func main() {

	defer func() {
		fmt.Println("Defer")
	}()

	lifoSimpleDefer()
}
