package oj

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
