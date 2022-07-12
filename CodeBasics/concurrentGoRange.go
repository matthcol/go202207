package main

import (
	"fmt"
	"time"
)

func fc3(c chan int, n int) {
	for x := 0; x < n; x++ {
		y := 2*x + 1
		fmt.Println("Inside f:", y)
		c <- y
	}
	close(c)
}

func testGoMethodsRange() {
	n := 5  // nombre appel
	cn := 1 // taille du chan
	c := make(chan int, cn)

	fmt.Println("Call f with n = ", n)
	go fc3(c, n)

	fmt.Println("After call f")

	s := 0
	for r := range c {
		s += r
		fmt.Println("Receive result:", r)
		time.Sleep((1 * time.Second))
	}
	fmt.Println("Sum:", s)
	extraResult, ok := <-c // detection de deadlock
	fmt.Println("Extra:", extraResult, ok)
} // here execute defer calls

func main() {
	testGoMethodsRange()
	fmt.Println("Main after test")

}
