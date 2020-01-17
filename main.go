package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	// testCanJump()

	testLP()

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
