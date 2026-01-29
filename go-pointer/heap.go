package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() *person { // nghĩa là sẽ trả về 1 địa chỉ của biến có kdl là person
	m := person{
		name: "John",
		age:  30,
	}
	fmt.Printf("%p\n", &m)
	return &m
}



func main1() {
	fmt.Printf("%p\n", initPerson())
}
