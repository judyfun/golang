package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{}
	p.Age = 20
	p.Name = "jack"
	fmt.Println("raw -> ", p)
	updateAge(p)
	fmt.Println("suffix -> ", p)

}

/**
1, 此时传递的是person的值拷贝的信息，并不是原始的指针地址
在函数中更改的值，并不能更新原始地址的值
*/
func updateAge(per person) {
	per.Age = 30
	fmt.Println("updateAge", per)
	fmt.Println("%p\n", per)
}
