/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

package main

import "fmt"

type strTestCase struct {
	haystack string
	needle   string
	result   int
}

func testStrStr() {
	cases := []strTestCase{
		strTestCase{"hello", "ll", 2},
		strTestCase{"aaaaa", "bba", -1},
		strTestCase{"aaaaa", "", 0},
		strTestCase{"", "bba", -1},
		strTestCase{"aaa", "aaaa", -1},
		strTestCase{"mississippi", "mississippi", 0},
		strTestCase{"mississippi", "issip", 4},
		strTestCase{"mississippi", "issipi", -1},
	}
	for _, v := range cases {
		result := strStr(v.haystack, v.needle)
		fmt.Printf("strStr(%v, %v)=%v, %v\n", v.haystack, v.needle, result, result == v.result)
	}
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	lneed := len(needle)
	lhay := len(haystack)
	if lhay < lneed {
		return -1
	}

	for i := 0; i < lhay; i++ {
		index := i
		for j := 0; j < lneed; j++ {
			if haystack[index] != needle[j] {
				break
			} else if j == lneed-1 {
				return index - lneed + 1
			}
			index++
			if index >= lhay {
				return -1
			}
		}
	}
	return -1
}