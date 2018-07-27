/*
接雨水
*/
package main

import "fmt"

type trapTestCase struct {
	v []int
	r int
}

func trapTest() {
	cases := []trapTestCase{
		trapTestCase{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		trapTestCase{[]int{4, 2, 3}, 1},
		trapTestCase{[]int{4, 9, 4, 5, 3, 2}, 1},
		trapTestCase{[]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}, 23},
		trapTestCase{[]int{5, 2, 1, 2, 1, 5}, 14},
	}
	for _, v := range cases {
		fmt.Println("test", trap(v.v), v.r)
	}
}

func trap(height []int) int {
	// fmt.Println(height)
	count := 0
	for i := 0; i < len(height); i++ {
		next, max := find(height[i:])
		// fmt.Printf("i is:%v ,next is:%v,max is %v\n", i, next, max)
		if next > 0 && max > 0 {
			for j := i; j < i+next; j++ {
				if height[j] < max {
					count += max - height[j]
					// fmt.Println("count+=", max-height[j])
				}
			}
			i += next - 1
		}

	}
	return count
}

func find(t []int) (int, int) {
	min := t[0]
	max := 0
	index := 0
	l := len(t)
	for i := 1; i < l; i++ {
		// if max == t[i-1] && t[i] < t[i-1] {
		// 	return i - 1, max
		// }
		if t[i] < min {
			min = t[i]
		} else if t[i] < t[0] && t[i] > max {
			max = t[i]
			index = i
		} else if t[i] >= t[0] {
			if min < t[0] {
				max = t[0]
				return i, max
			}
			return -1, -1
		}
	}
	// if min < t[0] && t[l-1] > min {
	// 	return l, max
	// }
	if min < t[0] && max > min {
		return index, max
	}
	return -1, -1
}
