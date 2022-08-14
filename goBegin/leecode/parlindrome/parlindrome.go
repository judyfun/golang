package parlindrome

import "fmt"

//https://leetcode.com/problems/palindrome-number/

func Expose(x int) {
	isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0, 10)
	numberToArray(x, nums)

	fmt.Println(nums)
	return false
}

func numberToArray(x int, nums []int) {
	if x == 0 {
		return
	} else {
		remainder := x % 10
		nums = append(nums, remainder)
	}

	numberToArray(x/10, nums)
}
