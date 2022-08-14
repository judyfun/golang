package two_sum

import "testing"

func TestTwoSum(t *testing.T) {

	inputNums1 := []int{0, 4, 3, 0}
	expectedNums1 := [2]int{1, 2}

	inputNums2 := []int{2, 7, 11, 15}
	expectedNums2 := [2]int{0, 1}
	testTable := []struct {
		nums     []int
		target   int
		expected [2]int
	}{{nums: inputNums1,
		target:   7,
		expected: expectedNums1},

		{nums: inputNums2,
			target:   9,
			expected: expectedNums2},
	}

	for _, tt := range testTable {
		actual := twoSum(tt.nums, tt.target)
		if actual != tt.expected {
			t.Errorf("twosum(%v,%v), expected: %v, actual: %v", tt.nums, tt.target, tt.expected, actual)
		}
	}
}
