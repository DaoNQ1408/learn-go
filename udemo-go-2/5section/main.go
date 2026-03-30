package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrValueTooLarge = errors.New("a is too large")

type OpError struct {
	Op      string
	Code    int
	Message string
}

// Implement the error interface for OpError
func (op OpError) Error() string {
	return op.Message
}

func newOpError(op string, code int, message string) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &OpError{
			Op:      "divide",
			Code:    500,
			Message: "division by zero",
		}
	}

	if a > 1000 {
		return 0, ErrValueTooLarge
	}
	return a / b, nil

}

func DoSomething() error {
	return newOpError("DoSomething", 500, "something went wrong")
}

func main() {
	value, err := divide(101, 0)
	if err != nil {
		var opErr *OpError
		// dùng As để ktra kdl của lỗi
		if errors.As(err, &opErr) {
			fmt.Printf("Operation: %s, Code: %d, Message: %s\n", opErr.Op, opErr.Code, opErr.Message)
		}
		// dùng Is để ktra lỗi có phải là ErrValueTooLarge hay không
		if errors.Is(err, ErrValueTooLarge) {
			fmt.Println(ErrValueTooLarge)
		}
	}
	fmt.Println(value)
}
