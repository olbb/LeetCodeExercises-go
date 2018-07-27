/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/2^2 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
*/

package main

import "fmt"

type powTestCase struct {
	x float64
	n int
	r float64
}

func testPow() {
	cases := []powTestCase{
		powTestCase{2, 10, 1024},
		powTestCase{2.1, 3, 9.261},
		powTestCase{2.00000, -2, 0.25000},
		powTestCase{0.00001, 2147483647, 0},
		powTestCase{0.44521, 0, 1},
	}
	for _, v := range cases {
		result := myPow(v.x, v.n)
		fmt.Printf("myPow(%v,%v) result=%v,%v\n", v.x, v.n, result, result == v.r)
	}
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	// count := 0
	// var result float64 = 1
	result := x
	for i := n - 1; i > 0; i-- {

		if i%2 == 0 {
			x = x * x
			i = i / 2
		}
		result *= x
		// count++
	}
	// fmt.Println("共循环", count)
	return result
}
