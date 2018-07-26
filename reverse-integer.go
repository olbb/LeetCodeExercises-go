/*
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	max := 0xffffffff / 2
	min := -0xffffffff + 1

	negative := x < 0
	if negative {
		x = -x
	}
	str := fmt.Sprintf("%d", x)
	// fmt.Println(str)
	var b strings.Builder
	for i := len(str); i > 0; i-- {
		s := str[i-1 : i]
		b.WriteString(s)
	}
	result, _ := strconv.Atoi(b.String())
	if negative {
		result = -result
	}
	if result > max || result < min {
		return 0
	}
	// fmt.Println(b.String())
	return result
}
