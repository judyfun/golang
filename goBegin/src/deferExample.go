package src

import (
	"fmt"
)

func main() {
	/*	A()
		B()
		C()*/

	homeWork1()
}

func homeWork1() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}
}

func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")

		}
	}()
	panic("panic in B")
	fmt.Println("Func B")
}

func C() {
	fmt.Println("Func C")
}
