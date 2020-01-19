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
		lpTestCase{"babad", "bab"},
		lpTestCase{"cbbd", "bb"},
		lpTestCase{"cbababd", "babab"},
		lpTestCase{"a", "a"},
		lpTestCase{"ac", "a"},
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
	// fmt.Println("deal", newStr)
	strs2 := strings.Split(newStr, "")
	l = len(strs2)
	var longest = strs2[0]
	longestl := 0
	longestr := 0
	for i := 1; i < l-1; i++ {
		// fmt.Println("index is ", i, ",len is:", l)
		curSize := 0
		// var tmp strings.Builder
		left := i
		right := i
		for {
			if left < 1 || right >= l-1 || strs2[left-1] != strs2[right+1] {
				if left%2 != 0 {
					left++

				}
				break
			}
			curSize = curSize + 1
			left = i - curSize
			right = i + curSize
		}
		// fmt.Printf("test left:%v,right:%v\n", left, right)
		ll := (right-left)/2 + 1
		ol := (longestr-longestl)/2 + 1
		if ll > ol {
			longestl = left
			longestr = right
		}

	}
	var tmp strings.Builder
	// fmt.Printf("result:%v, left is :%v ,ll:%v\n", strs2[left:right+1], left, ll)
	for j := longestl; j <= longestr; j = j + 2 {
		tmp.WriteString(strs2[j])
	}
	longest = tmp.String()
	return longest
}
