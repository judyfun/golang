package main

import "fmt"

func main() {
	//breakDemo()
	//gotoDemo()
	//contiueDemo()
	demo()
}

func demo() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL
		}
	}
	fmt.Println("ok")
}

func contiueDemo() {
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			// 进行下一次循环，next
			continue LABEL

		}
	}

	fmt.Println("ok")

}

func gotoDemo() {
	for {
		for i := 0; i < 10; i++ {
			if i == 5 {
				fmt.Println(i)
				// 调整执行位置
				goto LABEL
			}
		}
	}
LABEL:
	fmt.Println("ok")

}

func breakDemo() {
LABEL:
	for {
		for i := 0; i < 10; i++ {
			break LABEL
		}
	}

	fmt.Println("ok")
	fmt.Println("----")
}
