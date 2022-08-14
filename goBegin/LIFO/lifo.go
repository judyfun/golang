package LIFO

import "fmt"

func Expose() {
	lifoSimple()
}

func lifoSimple() {
	var stack []string
	stack = append(stack, "Hello")
	stack = append(stack, "World")

	for len(stack) > 0 {
		n := len(stack) - 1
		fmt.Println(stack[n])

		stack = stack[:n]
	}
}
