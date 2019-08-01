package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b, c, d := 1, 2, 3, 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	e, _, f, g := 5, 6, 7, 8
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var z int = 65
	y := string(z)
	fmt.Println(y)

	x := strconv.Itoa(z)
	fmt.Println(x)
	i, i2 := strconv.Atoi(x)
	fmt.Println(i, i2)

	fmt.Println(1 ^ 2)

	u := 1
	var p *int = &u
	fmt.Println(*p)

}
