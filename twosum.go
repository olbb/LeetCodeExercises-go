/* 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1] */

package main

import "fmt"

func twosum() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(getnum(nums, 13))
}

func getnum(nums []int, t int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i != j && num1+num2 == t {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}
