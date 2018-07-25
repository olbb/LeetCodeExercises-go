/*
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
*/

package main

func lengthOfLongestSubstring(s string) int {
	longest := ""
	temps := make(map[string]int)
	index := 0
	strLen := len(s)
	for j := 0; j < strLen; j++ {
		key := s[j : j+1]

		_, find := temps[key]

		if find { //找到了 移动index重新找
			index++
			temps = make(map[string]int)
			j = index - 1
		} else {
			l := j - index + 1 //当前处理的字符串长度
			if l > len(longest) {
				longest = s[index : j+1]
			}
			temps[key] = j
		}

	}
	return len(longest)

}
