package slice

import "fmt"

func Expose() {
	appendOneSliceToAnother()
}

// appendToAnEmptySlice this fun has a risk
func appendToAnEmptySlice() {
	var a []int
	a = append(a, 3, 4)
	fmt.Println(a)
}

func appendOneSliceToAnother() {
	a := []int{3, 4}
	b := []int{5, 6}
	a = append(a, b...)
	fmt.Println(a)
}
