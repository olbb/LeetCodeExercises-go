/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

package main

import "fmt"

type canJumpTestCase struct {
	nums   []int
	result bool
}

func testCanJump() {
	cases := []canJumpTestCase{
		// canJumpTestCase{[]int{2, 3, 1, 1, 4}, true},
		// canJumpTestCase{[]int{1, 2, 1}, true},
		// canJumpTestCase{[]int{}, true},
		// canJumpTestCase{[]int{0, 1}, false},
		canJumpTestCase{[]int{1, 2}, true},
		// canJumpTestCase{[]int{3, 2, 1, 0, 4}, false},
	}
	for _, v := range cases {
		result := canJump(v.nums)
		fmt.Printf("canJump %v, result %v, %v\n", v.nums, result, v.result == result)
	}
}

func canJump(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}

	count := 0
	for i := l - 1; i > 0; {
		index := getMaxLengthIndex2(nums[:i])
		if index >= 0 {
			i = index
		} else {
			return false
		}
		count++
	}
	return true
}

func getMaxLengthIndex2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] > 0 {
			return 0
		}
		return -1
	}

	maxStep := -1
	maxStepIndex := -1
	for i := l - 1; i >= 0; i-- {
		if nums[i] >= l-i { //可以跳跃
			if l-i >= maxStep {
				maxStep = l - i
				maxStepIndex = i
			}
		}
	}
	// fmt.Printf("found index:%v, value:%v\n", maxStepIndex, maxStep)
	return maxStepIndex
}
