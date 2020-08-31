package leetcode

func TwoSum(nums []int, target int) []int {
	index := map[int]int{}

	for i, num := range nums {
		diff := target - num
		if j, ok := index[diff]; ok {
			return []int{j, i}
		}
		index[num] = i
	}
	return nil
}

func TwoSumV2(nums []int, target int) []int {
	index := map[int]bool{}

	for _, num := range nums {
		diff := target - num
		if _, ok := index[diff]; ok {
			if num > diff {
				return []int{diff, num}
			}
			return []int{num, diff}
		}
		index[num] = true
	}
	return nil
}
