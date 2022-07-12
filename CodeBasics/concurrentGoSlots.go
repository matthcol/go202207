package main

import (
	"fmt"
	"time"
)

func fc2(c chan int, n int) {
	for x := 0; x < n; x++ {
		y := 2*x + 1
		fmt.Println("Inside f:", y)
		c <- y
	}
}

func testGoMethodsSlots() {
	n := 5  // nombre appel
	cn := 5 // taille du chan
	c := make(chan int, cn)

	fmt.Println("Call f with n = ", n)
	go fc2(c, n)

	fmt.Println("After call f")

	s := 0
	for x := 0; x < n; x++ {
		r := <-c
		s += r
		fmt.Println("Receive result:", r)
		time.Sleep((2 * time.Second))
	}
	fmt.Println("Sum:", s)
	// <-c // detection de deadlock

} // here execute defer calls

func main() {
	testGoMethodsSlots()
	fmt.Println("Main after test")

}
