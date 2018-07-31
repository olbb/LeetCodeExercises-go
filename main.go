package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	testCanJump()

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
