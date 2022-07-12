package main

import "fmt"

// exercise from Tour of Go : https://go.dev/tour/moretypes/26

type IntGenerator func() int

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() IntGenerator {
	a := 0
	b := 0
	return func() int {
		if b == 0 {
			b = 1
		} else {
			a, b = b, a+b
		}
		return a
	}
}

func main() {
	var f IntGenerator = fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
