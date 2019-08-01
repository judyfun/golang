package main

import "fmt"

func main() {
	a := 1
	for {
		a++
		if a > 4 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("----")

	for a < 9  {
		a++
		fmt.Println(a)
	}

	fmt.Println("----")

	for i := 0; i<4; i++  {
		fmt.Println(i)
	}
	fmt.Println("Over")

}
