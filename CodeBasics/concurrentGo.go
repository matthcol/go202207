package main

import (
	"fmt"
	"time"
)

func fc(c chan int, x int) {
	y := 2*x + 1
	time.Sleep((7 - time.Duration(x)) * time.Second)
	fmt.Println("Inside f:", y)
	c <- y
}

func testGoMethods() {
	c := make(chan int)
	for x := 0; x < 5; x++ {
		fmt.Println("Call f with x = ", x)
		go fc(c, x)
	}
	fmt.Println("After Loop")
	// time.Sleep(10 * time.Second) // to wait for all goroutines to end

	// join goroutine : 5 x receive
	// + calcul on results
	s := 0
	for x := 0; x < 5; x++ {
		r := <-c
		s += r
		fmt.Println("Receive result:", r)
	}
	fmt.Println("Sum:", s)
	// <-c // detection de deadlock

} // here execute defer calls

func main() {
	testGoMethods()
	fmt.Println("Main after test")

}
