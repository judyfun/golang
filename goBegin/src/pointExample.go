package main

import "fmt"

/**
&符号的意思: 对变量取地址，如：变量a的地址是&a
*符号的意思: 对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了

*和 & 可以互相抵消,同时注意，*&可以抵消掉，但&*是不可以抵消的
 */
func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b

	fmt.Println("a = ",a)
	fmt.Println("&a = ",&a)
	fmt.Println("*&a = ",*&a)
	fmt.Println("b = ",b)
	fmt.Println("&b = ",&b)
	fmt.Println("*&b = ",*&b)
	fmt.Println("*b = ",*b)
	fmt.Println("c = ",c)
	fmt.Println("*c = ",*c)
	fmt.Println("&c = ",&c)
	fmt.Println("*&c = ",*&c)
	fmt.Println("**c = ",**c)
	fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)
	fmt.Println("x = ",x)
}
