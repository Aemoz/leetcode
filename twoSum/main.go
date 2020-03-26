package main

import "fmt"

func twoSum(nums []int, target int) []int {

	for i := range nums {
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if (a + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nums
}

func main() {
	//
	nums := []int{3, 2, 4}
	target := 6
	xx := twoSum(nums, target)
	fmt.Print(xx)

}
