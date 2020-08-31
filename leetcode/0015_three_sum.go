package leetcode

func ThreeSum(nums []int) [][]int {

	result := [][]int{}
	for i, num := range nums {
		t := TwoSumV2(nums[i+1:], -num)
		if t != nil {
			t = append([]int{num}, t...)
			result = append(result, t)
		}
	}
	return result
}
