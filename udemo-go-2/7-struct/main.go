package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Software Engineer",
		Salary:    80000,
		IsActive:  true,
		JoinedAt:  time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC), // time.Now()
	}

	fmt.Printf("%+v\n", jane)

	janPtr := &jane
	janPtr.Salary = 85000
	janPtr.IsActive = false
	fmt.Printf("%+v\n", jane)
	fmt.Println(&janPtr) // jane address
	fmt.Println(*janPtr) // jane value
	fmt.Println(&jane)   // & + jane value
}
