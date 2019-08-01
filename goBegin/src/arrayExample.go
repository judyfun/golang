package main

import "fmt"
import "strconv"

func main() {
	//initExample()
	//init2Example()
	//init3()
	//init4()
	//init5()
	//piontArr()
	//piontArr2()
	//compare()
	//newArr()
	//setArr()
	//muliArr()
	pop()
}

// 冒泡排序
func pop() {
	a := [...]int{4, 3, 5, 2, 9, 6, 8}
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			fmt.Println("i - " + strconv.Itoa(i))
			fmt.Println("j - " + strconv.Itoa(j))
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
			fmt.Println(a)
		}
	}
	fmt.Println(a)

}

// 多维数组
func muliArr() {
	a := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(a)

	b := [2][3]int{
		{1: 3},
		{2: 1}}
	fmt.Println(b)

	c := [...][3]int{
		{1: 3},
		{2: 1}}
	fmt.Println(c)
}

//给数组赋值,普通数组，和指向数组的指针，都可以通过下表操作数组中的值
func setArr() {
	a := [10]int{}
	a[1] = 2
	fmt.Println(a)

	b := new([10]int)
	b[1] = 2
	fmt.Println(b)
	fmt.Println(*b)

}

// new关键字，返回地址
func newArr() {
	p := new([10]int)
	fmt.Println(p)
	fmt.Println(*p)

}

// 数组比较，必须相同类型
func compare() {
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	//c := [1]int{1}
	fmt.Println(a == b)
	//fmt.Println(a == c)

}

func piontArr2() {
	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)
}

// 数组指针
func piontArr() {
	a := [...]int{9: 1}
	var p *[10]int = &a
	fmt.Println(p)

}

func init5() {
	a := [...]int{0: 1, 8: 4, 9: 9}
	fmt.Println(a)

}

func init4() {
	a := [...]int{1, 2, 4, 5}
	fmt.Println(a)

}

//初始化，序号第无的值是4，第九的值为1，序号从零开始算
func init3() {
	a := [10]int{5: 4, 9: 1}
	fmt.Println(a)

}

// 数组初始化2
func init2Example() {
	var a = [2]int{1, 1}
	b := [2]int{1}
	fmt.Println(a)
	fmt.Println(b)
}

//数组如果没有初始化，那么就会默认值，"零值"
func initExample() {
	var a [2]int
	var b [10]string

	fmt.Println(a)
	fmt.Println(b)
}
