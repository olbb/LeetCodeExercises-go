/*
初始时有 n 个灯泡关闭。 第 1 轮，你打开所有的灯泡。 第 2 轮，每两个灯泡你关闭一次。 第 3 轮，每三个灯泡切换一次开关（如果关闭则开启，如果开启则关闭）。第 i 轮，每 i 个灯泡切换一次开关。 对于第 n 轮，你只切换最后一个灯泡的开关。 找出 n 轮后有多少个亮着的灯泡。

示例:

输入: 3
输出: 1
解释:
初始时, 灯泡状态 [关闭, 关闭, 关闭].
第一轮后, 灯泡状态 [开启, 开启, 开启].
第二轮后, 灯泡状态 [开启, 关闭, 开启].
第三轮后, 灯泡状态 [开启, 关闭, 关闭].

你应该返回 1，因为只有一个灯泡还亮着。

1.解法1
[0 1 1 2 1 3 1 3 2 3 1 5 1 3 3 4 1 5 1 5 3 3 1 7 2 3 3 5 1 7 1 5 3 3 3 8 1 3 3 7 1 7 1 5 5 3 1 9 2 5 3 5 1 7 3 7 3 3 1 11 1 3 5 6 3 7 1 5 3 7 1 11 1 3 5 5 3 7 1 9]
输出结果可以转化为求每一位约数个数的题

2.解法2
下面我来解释一下。对于每个数字来说，除了平方数，都有偶数个因数。

如6有4个因数：1×6=6，2×3=6

如60有个因数：1×60=60，2×30=60，3×20=60，4×15=60，5×12=60，6×10=60

可以看出，非平方数的因数总是成对出现的，只有平方数的因数才是奇数，因为平方数除平方根外，其他的因数都是成对出现的！对于当前的开关灯泡问题，
可知到最后处在平方数位置的灯泡一定是开启的，其他位置的灯泡一定是关闭的。而要计算一个数之下有多少小于或等于它的平方数，使用一个开平方用的函数就可以了。



*/

package main

import (
	"fmt"
	"math"
)

func bulbTest() {
	fmt.Println(bulbSwitch(99999990))
	// getPrimes(9999999)
	// fmt.Println(t)
}

// func getPrimes(n int) int {
// 	temps := make(map[int]int)
// 	s := math.Sqrt(float64(n))
// 	for i := 0; i < int(s); i++ {
// 		if n % i == 0 {
// 			temps[]
// 		}
// 	}
// }

//太慢
func getPrimes(n int) int {
	temps := make(map[int]int)
	for i := 2; i < n; i++ {
		for i != n && n%i == 0 {
			temps[i] = temps[i] + 1
			n = n / i
		}
	}
	temps[n] = temps[n] + 1
	fmt.Println(temps)
	count := 1
	for _, v := range temps {
		count *= v + 1
	}
	return count
}

var c int

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func bulbSwitchOld(n int) int {
	if n == 0 {
		return 0
	}
	nums := make([]bool, n)
	// accessTimes := make([]bool, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			for i := 0; i < n; i++ {
				nums[i] = true
			}
		} else {
			changeNums(nums, i+1)
		}
		// fmt.Println(nums)
	}
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] {
			count++
		}
	}
	fmt.Println("运算", c)
	count2 := 0
	// for _, v := range accessTimes {
	// 	if v%2 == 0 {
	// 		count2++
	// 	}
	// }

	// fmt.Println(accessTimes)
	fmt.Println(count2)
	return count
}

func changeNums(nums []bool, times int) {
	l := len(nums)
	for i := times - 1; i < l; i += times {
		nums[i] = !nums[i]
		c++
	}
}

func change(status int) int {
	if status == 1 {
		return 0
	}
	return 1
}
