package src

import (
	"fmt"
)

func main() {
	//
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")

	}

	fmt.Println("--------")

	// 表达式
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")

	}

	fmt.Println("--------")

	// 初始化
	switch b := 1; {
	case b >= 0:
		fmt.Println("b>=0")
		fallthrough
	case b >= 1:
		fmt.Println("b>=1")
	default:
		fmt.Println("None")
	}

}
