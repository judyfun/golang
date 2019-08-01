package main

import "fmt"

func main() {
	initExample()
}

//数组如果没有初始化，那么就会默认值，"零值"
func initExample() {
	var a [2]int
	var b [10]string

	fmt.Println(a)
	fmt.Println(b)
}
