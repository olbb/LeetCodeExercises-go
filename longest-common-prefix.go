/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/
package main

import (
	"fmt"
)

type lCPTestCase struct {
	strs []string
	r    string
}

func testLCP() {
	cases := []lCPTestCase{
		lCPTestCase{[]string{"flower", "flow", "flight"}, "fl"},
		lCPTestCase{[]string{"dog", "racecar", "car"}, ""},
		lCPTestCase{[]string{}, ""},
	}
	for _, v := range cases {
		r := longestCommonPrefix(v.strs)
		fmt.Printf("longestCommonPrefix(strs %v) = %v , %v\n", v.strs, r, v.r == r)
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	index := -1
	l := len(strs[0])

	// for _, v := range strs {
	// 	if v == "" {
	// 		return ""
	// 	}
	// 	l = minInt(l, len(v))
	// }
	for i := 0; i < l; i++ {
		b := strs[0][i : i+1]
		// fmt.Println("i:", i, " b:", b)
		for _, v := range strs {
			if i >= len(v) || b != v[i:i+1] {
				// fmt.Println("breaked,index is ", index)
				return strs[0][:index+1]
			}
		}
		index = i
	}
	return strs[0][:index+1]
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
