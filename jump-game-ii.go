/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
说明:

假设你总是可以到达数组的最后一个位置。
*/
package main

import "fmt"

type jumpTestCase struct {
	nums   []int
	result int
}

func testJump() {
	cases := []jumpTestCase{
		jumpTestCase{[]int{2, 3, 1, 1, 4}, 2},
		jumpTestCase{[]int{1, 2, 1}, 2},
		jumpTestCase{[]int{}, 0},
	}
	for _, v := range cases {
		result := jump(v.nums)
		fmt.Printf("jump %v, result %v, %v\n", v.nums, result, v.result == result)
	}
}

func jump(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	if l == 2 {
		return 1
	}
	count := 0
	for i := l - 1; i > 0; {
		index := getMaxLengthIndex(nums[:i])
		if i >= 0 {
			i = index
		} else {
			return -1
		}
		count++
	}
	return count
}

func getMaxLengthIndex(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		return 0
	}

	maxStep := -1
	maxStepIndex := 0
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
