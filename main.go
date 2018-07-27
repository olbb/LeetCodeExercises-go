package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	testPow()

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
