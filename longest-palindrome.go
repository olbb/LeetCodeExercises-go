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

type lpTestCase struct {
	input  string
	result string
}

func testLP() {
	cases := []lpTestCase{
		// lpTestCase{"babad", "bab"},
		// lpTestCase{"cbbd", "bb"},
		// lpTestCase{"cbababd", "babab"},
		// lpTestCase{"a", "a"},
		// lpTestCase{"ac", "a"},
		lpTestCase{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
	}
	for _, v := range cases {
		r := longestPalindrome(v.input)
		fmt.Printf("longestPalindrome(strs %v) = %v , %v\n", v.input, r, v.result == r)
	}
}

func longestPalindrome(str string) string {
	var sb strings.Builder
	l := len(str)
	if l <= 1 {
		return str
	}
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
	var longest = strs2[0]
	for i := 1; i < l-1; i++ {
		// fmt.Println("index is ", i, ",len is:", l)
		curSize := 0
		// var tmp strings.Builder
		left := i
		right := i
		for {
			// tmp.Reset()
			curSize = curSize + 1
			left = i - curSize
			right = i + curSize
			if left < 0 || right >= l || strs2[left] != strs2[right] {
				break
			}
			// fmt.Println("find", strs2[left:right+1], "i=", i, left, right, "left%2 == 0", left%2 == 0)
			// if left%2 == 0 { // b#b形式

			// 	for j := left; j <= right; j = j + 2 {
			// 		// fmt.Println("append:" + strs2[j])
			// 		tmp.WriteString(strs2[j])
			// 	}
			// } else { // #b# 形式
			// 	for j := left + 1; j <= right; j = j + 2 {
			// 		// fmt.Println("append:" + strs2[j])
			// 		tmp.WriteString(strs2[j])
			// 	}
			// }
			if left%2 != 0 { //  #b#
				left++
			}

			//strnew := tmp.String()
			// fmt.Println("result:" + strnew)
			// if len(strnew) > 0 {
			// 	// fmt.Println("got :", strnew)
			// 	if len(longest) < len(strnew) {
			// 		longest = strnew
			// 	}
			// }
		}
		ll := (right - left)/2
		if (ll > len(longest)) {
			for (j == left; j<= right; j = j + 2) {
				
			}
		}

	}
	return longest
}
