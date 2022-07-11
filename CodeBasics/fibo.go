// suite valeur 0 1 1 2 3 5 ...
// stored in an slice of n values
package main

import "fmt"

func fibo(n int) []uint64 {
	result := make([]uint64, n)
	// todo : protect 2 assignements if n = 0 or 1
	result[0] = 0
	result[1] = 1
	for i := 0; i < n-2; i++ {
		// for i := range result[2:] {
		result[i+2] = result[i] + result[i+1]
	}
	return result
}

func main() {
	f := fibo(10)
	fmt.Println(f)
}
