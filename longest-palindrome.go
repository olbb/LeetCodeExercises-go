package main

import (
	"fmt"
	"strings"
)

/**
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	示例 1：

	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。
	示例 2：

	输入: "cbbd"
	输出: "bb"
 **/

func longestPalindrome(str string) string {
	var sb strings.Builder
	l := len(str)
	strs := strings.Split(str, "")
	for i := 0; i < l; i++ {

		sb.WriteString(strs[i])
		if i < l-1 && i != l {
			sb.WriteString("#")
		}
	}
	newStr := sb.String()
	fmt.Println("deal", newStr)
	strs2 := strings.Split(newStr, "")
	l = len(strs2)
	var longest string
	for i := 1; i < l-1; i++ {
		curSize := 0
		var tmp strings.Builder
		for {
			curSize = curSize + 1
			left := i - curSize
			right := i + curSize
			if left < 0 || right > l || strs2[left] != strs2[right] {
				break
			}
			fmt.Println("find", strs2[left:right+1], "i=", i, left, right)
			if left%2 == 0 { // #b# 形式

				for j := left; j <= right; j = j + 2 {
					tmp.WriteString(strs2[j])
				}
			} else { // b#b形式
				for j := left + 1; j <= right; j = j + 2 {
					tmp.WriteString(strs2[j])
				}
			}
			fmt.Println("result:" + tmp.String())

		}

		strnew := tmp.String()
		if len(strnew) > 1 {
			fmt.Println("got :", strnew)
			if len(longest) < len(strnew) {
				longest = strnew
			}
		}

	}
	return ""

}
