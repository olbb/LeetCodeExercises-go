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

func longestCommonPrefix(strs []string) string {
	index := 0
	l := len(strs[0])
	for _, v := range strs {
		l = minInt(l, len(v))
	}
	for i := 0; i < l; i++ {
		b := strs[0][i : i+1]
		for _, v := range strs {
			if b != v[i:i+1] {
				break
			}
		}
	}
	return strs[0][:index]
}

func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
