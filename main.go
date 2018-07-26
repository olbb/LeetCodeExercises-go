package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	fmt.Println(reverse(1534236469))

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
