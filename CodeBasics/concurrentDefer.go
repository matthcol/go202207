package main

import (
	"fmt"
	"time"
)

func f(x int) {
	y := 2*x + 1
	time.Sleep(2 * time.Second)
	fmt.Println("Inside f:", y)
}

func testDeferreMethods() {
	for x := 0; x < 5; x++ {
		fmt.Println("Call f with x = ", x)
		defer f(x)
	}
	fmt.Println("After Loop")
} // here execute defer calls

func main() {
	testDeferreMethods()
	fmt.Println("Main after test")
}
