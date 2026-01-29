package main

import (
	"fmt"
)

func main() {
	a := 10
	// b := "hello"
	fmt.Println(&a, a)
	power(&a) // truyền vào địa chỉ của a
	fmt.Println("a is here")
	fmt.Println(&a, a)
	main1()
}

func power(v *int) { // lấy địa chỉ đc truyền vào
	*v *= *v // update value tại địa chỉ đó
	fmt.Println("v is here")
	fmt.Println(*v, v)
	// giá trị mới và địa chỉ đc truyền vào
}

// &a: địa chỉ của biến a
// v = &a : v là địa chỉ của a
// *v: giá trị của biến tại địa chỉ được lưu trong v, tức là giá trị của a.

