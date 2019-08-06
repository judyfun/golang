package main

import (
	"fmt"
)

func main() {
	//b("hello", 1, 2, 3)
	/*	a := 1
		b := 2
		basicType(a, b)
		fmt.Println("----raw")
		fmt.Println(a)
		fmt.Println(b)*/

	/*	s1 := []int{1, 2, 3}
		directSliceType(s1)
		fmt.Println(s1)*/

	/*a1 := 1
	fmt.Println("raw address -> ", &a1)
	basicTypeByPoint(&a1)
	fmt.Println("after value -> ", a1)*/

	/*a := funcType
	a()*/

	//anonymousFunc()
	in := Increase()
	fmt.Println(in())
	fmt.Println(in())

}

/**
闭包作为函数的返回值
 */
func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}


func anonymousFunc() {
	var f = func(int) {}
	f = func(i int) {
		fmt.Println(i)
	}
	f(2)

	f = func(i int) {
		fmt.Println(i * i * i)
	}
	f(2)
}

/**
1,函数类型
函数可以作为一种类型，赋值给变量
*/
func funcType() {
	fmt.Println("funcType")
}

/**
1，基本类型的引用传递
将基本类型的内存地址作为参数，传递
函数对内存地址直接新型操作，则直接会改变原来的值
*/
func basicTypeByPoint(a *int) {
	fmt.Println("a address -> ", a)
	*a = 2
	fmt.Println("a value -> ", *a)
}

/**
1,直接传递一个slice对象作为参数，会将地址作为拷贝传递过去
如果在函数方法中，对slice进行操作，或者改变，则会影响原来的slice的值
也就是对原来的slice直接进行操作，因为传递的是地址的拷贝
*/
func directSliceType(s []int) {
	s[0] = 100
	s[0] = 101
	s[0] = 102
	fmt.Println(s)
}

/**
1,基本类型的值传递，其实传递的是值拷贝，给函数的参数
在函数的方法中，改变值，并不影响原来地址中的值

*/
func basicType(s ...int) {
	s[0] = 99
	s[1] = 100
	fmt.Println(s)
}

/**
1,可变参数
	可变参数用 ... 代替，接收的是参数数据类型是 slince
	可变参数，必须放在参数列表的最后一个参数
*/
func b(b string, s ...int) {
	fmt.Print(s)
}

/**
1,多参数，多返回值
*/
func a(a int, b string) (c int, d string) {
	return
}
