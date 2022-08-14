package two_sum

import (
	"fmt"
)

//https://leetcode.com/problems/two-sum/

func main() {
	//nums := []int{6, 2, 8, 7, 11, 15}
	nums1 := []int{0, 4, 3, 0}
	target := 0
	result := twoSum(nums1, target)
	fmt.Printf("result: i= %v, j=%v \n", result[0], result[1])
}

func twoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}
