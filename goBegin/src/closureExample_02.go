package src

func main() {
	/*i := incr()

	println("%p", i)
	println(i()) // 1
	println(i()) // 2
	println(i()) // 3

	k := incr()
	println(k())
	println("%p", k)*/

	// 这里调用了三次 incr()，返回了三个闭包，这三个闭包引用着三个不同的 x，它们的状态是各自独立的。
		println(incr()()) //1
		println(incr()()) //1
		println(incr()()) //1

	demo1()
}

func cEDemo2()  {
	x := 1
	func() {
		println(&x) // 0xc0000de790
	}()
	println(&x) // 0xc0000de790

}

/**
因为闭包对外层词法域变量是引用的，所以这段代码会输出 3。

可以想象 f 中保存着 x 的地址，它使用 x 时会直接解引用，所以 x 的值改变了会导致 f 解引用得到的值也会改变。
*/
func demo1() {
	x := 1
	f := func() {
		println(&x, x)
	}
	x = 2
	x = 3
	println(&x, x)
	f() // 3
}

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
