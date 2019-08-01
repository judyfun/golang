package main

import (
	std "fmt"
	"math"
)

const (
	name = "aa"
	pass = "123"
)

var (
	plat int
	cap  string
)

type (
	phone int
	paper string
)

func main() {
	std.Println("Hello 世界")
	std.Println("hello,%s\n", "world")
	std.Println("Sqart(9)=%f", math.Sqrt(9.0))
	std.Println(name)
	std.Println(plat)
	std.Println(math.MinInt8)

	b := false
	std.Println(b)
}
