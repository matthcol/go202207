package main

import "fmt"

type City struct {
	Name string
	Pop  uint
}

// Display as a function
func DisplayF(city City) {
	fmt.Printf("%s (pop: %d)\n", city.Name, city.Pop)
}

// Display as a method for type City
// NB1: (city City) is called receiver
// NB2: receiver is a copy
func (city City) Display() {
	fmt.Printf("%s (pop: %d)\n", city.Name, city.Pop)
}

// method ToString
// NB: receiver can be passed as a pointer also
func (city City) ToString() string {
	return fmt.Sprintf("%s (pop: %d)", city.Name, city.Pop)
}

// receiver as a pointer
func (city *City) IncrPop(delta uint) uint {
	city.Pop += delta
	return city.Pop
}

func IncrPopF(city *City, delta uint) uint {
	city.Pop += delta
	return city.Pop
}

func structBasics() {
	fmt.Println("Struct simple examples")

	// City by default: Name empty string, Pop 0
	var city1 City
	fmt.Println(city1)

	city2 := City{"Lyon", 1500000}
	fmt.Println(city2)

	city3 := City{Name: "Toulouse", Pop: 1300000}
	fmt.Println(city3, city3.Name, city3.Pop)

	city4 := City{Name: "Pau"}
	fmt.Println(city4)

	city5 := City{}
	fmt.Println(city5)

	city3.Pop++
	fmt.Println(city3, city3.Name, city3.Pop)

	DisplayF(city3)
	city3.Display()
	fmt.Println(city3.ToString())

	// call with receiver as a pointer : implicit
	newPop := city3.IncrPop(50000)
	city3.Display()
	fmt.Println("New population:", newPop)

	newPop = IncrPopF(&city3, 25000)
	fmt.Println("New population (2):", newPop, city3.ToString())

	var cityPtr *City = &city2
	newPop = cityPtr.IncrPop(7)
	fmt.Println("New population (3):", newPop, cityPtr.ToString())

	var ptr **City
	ptr = &cityPtr
	// implicit method call is not a gift with 2 levels of pointers
	fmt.Println(ptr, (*ptr).ToString())
}

type MyInt int

func (x *MyInt) Plus(delta MyInt) MyInt {
	return (*x) + delta
}

func (x *MyInt) Incr(delta MyInt) {
	*x += delta
}

func (x *MyInt) IncrI(delta int) {
	*x += MyInt(delta)
}

// func (x *MyInt) Plus(delta *MyInt) MyInt {
// 	return (*x) + (*delta)
// }

// func (x MyInt) Plus(delta MyInt) MyInt {
// 	return x + delta
// }

func addFunctionality() {
	var x MyInt = 3
	var y MyInt = 5
	var i int = 100
	// z := x.Plus(&y) // with arg as a pointer
	z := x.Plus(12)
	fmt.Println(z)
	z = x.Plus(y)
	fmt.Println(z)
	z.Incr(7)
	fmt.Println(z)
	z.Incr(y)
	fmt.Println(z)
	z.Incr(MyInt(i))
	fmt.Println(z)
	z.IncrI(i)
	fmt.Println(z)
}

func main() {
	// structBasics()
	addFunctionality()
}
