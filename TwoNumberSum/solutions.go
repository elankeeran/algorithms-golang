package main

import (
	"sort"
)

func main() {}

//Solution1 If any two numbers in the input array sum
//   up to the target sum, the function should return them in an array, in any
//   order. If no two numbers sum up to the target sum, the function should return
//   an empty array.
func Solution1(target int, nums ...int) []int {
	for i := 0; i < len(nums)-1; i++ {
		firstNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondNum := nums[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

//Solution2 If any two numbers in the input array sum
//   up to the target sum, the function should return them in an array, in any
//   order. If no two numbers sum up to the target sum, the function should return
//   an empty array.
func Solution2(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num

		if found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}

//Solution3 If any two numbers in the input array sum
//   up to the target sum, the function should return them in an array, in any
//   order. If no two numbers sum up to the target sum, the function should return
//   an empty array.
func Solution3(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
