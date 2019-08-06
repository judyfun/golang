package main

import "fmt"

type human struct {
	Sex int
}

type people struct {
	Sex int
}

type teacher struct {
	human
	people
	Sex  int
	Name string
	Age  int
}

type student struct {
	human
	people
	Name string
	Age  int
}

func main() {
	initStruct()
}

/*
对象共有的类型，可以通过嵌套的方式存储在其他的struct中
默认的字段名称，就是类型名称 -> human human , 所以初始化的是字段名称是 human

1,如果只引用了外面一个strct type,那么，struct里面的字段，且不跟本struct字段名冲突，可以直接引用
2，如果只引用了一个，且字段名冲突，那么直接引用的是本身的字段
3，引用了多个struct，并且有重名的字段名，则规则如下面的例子，要带上标识的struct

*/
func initStruct() {
	t := teacher{Name: "jack", Age: 24, human: human{Sex: 1}}
	s := student{Name: "jack", Age: 13, human: human{Sex: 0}}

	fmt.Println("teacher -> ", t)
	fmt.Println("student -> ", s)

	t.Age = 1
	t.Sex = 100
	t.people.Sex = 101
	t.human.Sex = 102

	fmt.Println(t)
}
