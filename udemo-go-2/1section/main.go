package main

import "fmt"

const HOST = "localhost"

const (
	Sunday = iota + 1
	Monday
)

func main() {

	var greeting string // can be empty but not nil
	greeting = "Hello, world!"

	println(greeting)

	var count = 10
	println(count)
	count = 20
	println(count)

	const coun1 = 10
	println(coun1)

	email := "hello@gmail.com"
	fmt.Println(email)
	fmt.Printf("%T \n", true)

	fmt.Printf("%T \n",HOST)
	fmt.Println(Sunday)
}
