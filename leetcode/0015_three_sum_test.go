package leetcode

import (
	"testing"

	assert "gopkg.in/go-playground/assert.v1"
)

func TestThreeSum(t *testing.T) {

	numbers := []int{-1, 0, 1, 2, -1, -4}
	t.Log("input:", numbers)

	output := ThreeSum(numbers)
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, output)
}
