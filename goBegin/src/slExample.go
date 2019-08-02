package main

import "fmt"

func main() {
	copyDemo2()

}

func copyDemo2()  {
	s1 := []int{1,2,3,4,5}
	s2 := []int{6,7,8}
	copy(s2[1:2], s1[3:4])
	fmt.Println(s1,s2)

}

func copyDemo1()  {
	s1 := []int{1,2,3,4,5}
	s2 := []int{6,7,8}
	copy(s2, s1)
	fmt.Println(s1,s2)

}

func copyDemo()  {
	s1 := []int{1,2,3,4,5}
	s2 := []int{6,7,8}
	copy(s1, s2)
	fmt.Println(s1,s2)

}
func demo5() {
	a := []byte{'a', 'b', 'c', 'd', 'e'}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(string(s1), string(s2))

	// s2 ,append,超过cap值，此时不会指向a数组了，会重新申请新的地址
	s2 = append(s1,'h','k','k','k','k','1','2')
	// s1 和s2 中，c这个是元素，是指向的同一个地址，改变s1中该地址的值c-z，
	// 那么s2中指向该地址的值不会变，因为s2不指向a了，重新申请了地址
	s1[0] = 'z'
	fmt.Println(string(s1), string(s2))

}
/**
1，指向同一个地址，此时同时指向数组a
如果有多个slice指向同一个地址，
如果其中一个发生改变，那么其他的都会发生改变
因为他们指向的都是这个地址，具体地址中的内容，怎么变化，他们只需要展示出来就好了
*/
func demo4() {
	a := []byte{'a', 'b', 'c', 'd', 'e'}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(string(s1), string(s2))

	// s1 和s2 中，c这个是元素，是指向的同一个地址，改变s1中该地址的值c-z，
	// 那么s2中指向该地址的值也会变成z
	s1[0] = 'z'
	fmt.Println(string(s1), string(s2))

}

/**
append
如果没超过cap，则都是一个地址
如果超过cap，则会重新换个新地址，把原来的值拷贝过去
*/
func demo3() {
	s1 := make([]int, 3, 6)
	fmt.Print("%p\n", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println("%v %p", s1, s1)
	s1 = append(s1, 4, 5, 6)
	fmt.Println("%v %p", s1, s1)

}

/**
数组在内存中是一个连续的内存块
sa 指向 cde
sb 指向 de
*/

func demo2() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	// sa 的容量最大的指针，指向的是连续内存块的尾部
	fmt.Println(len(sa), cap(sa))
	fmt.Println(string(sa))
	sb := sa[1:3]
	fmt.Println(len(sb), cap(sb))
	fmt.Println(string(sb))

}
func demo1() {
	sl := make([] int, 3, 10)
	fmt.Println(len(sl), cap(sl), )
	fmt.Println(sl)
}
