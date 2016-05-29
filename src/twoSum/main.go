package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		for j, v2 := range nums[i+1:] {
			if (v + v2) == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

func intersect(nums1 []int, nums2 []int) []int {

}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}
