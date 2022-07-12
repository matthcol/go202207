package main

import "fmt"

type City2 struct {
	Name string
	Pop  uint
	F    func(p uint) uint
}

func main() {
	city := City2{Name: "Lyon", Pop: 1500000}
	city.F = func(x uint) uint { return x + 1 }
	fmt.Println(city)
	y := city.F(city.Pop)
	fmt.Println(y)
}
