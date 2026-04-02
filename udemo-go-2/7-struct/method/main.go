package main

type Person struct {
	Name string
	Age  int
}

func (p Person) updateName(newName string) { // dùng để tính toán, trả ra value mới, kh đụng gì đến struct gốc
	p.Name = newName
}

func (p *Person) updateName2(newName string) { // dùng để thay đổi giá trị của struct gốc thông qua pointer
	p.Name = newName
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	person.updateName("Bob")
	println(person.Name) // Output: Alice
	person.updateName2("Charlie")
	println(person.Name) // Output: Charlie
}
