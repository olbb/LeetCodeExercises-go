package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(isPalindrome(1001))

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
