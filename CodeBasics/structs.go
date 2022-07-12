package main

import "fmt"

type UintIncrementable interface {
	Incr(i uint) uint
	ToString() string
}

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

// implementation of interface Stringer
func (city City) String() string {
	return city.ToString()
}

// receiver as a pointer
func (city *City) Incr(delta uint) uint {
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
	newPop := city3.Incr(50000)
	city3.Display()
	fmt.Println("New population:", newPop)

	newPop = IncrPopF(&city3, 25000)
	fmt.Println("New population (2):", newPop, city3.ToString())

	var cityPtr *City = &city2
	newPop = cityPtr.Incr(7)
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

type MyUint uint

func (myuint *MyUint) Incr(i uint) uint {
	*myuint += MyUint(i)
	return uint(*myuint)
}

func (myuint *MyUint) ToString() string {
	return fmt.Sprintf("%d", *myuint)
}

func doSomethingWithUintIncrementable(data UintIncrementable) {
	data.Incr(111)
	fmt.Println(data.ToString())

	fmt.Printf("Data interface: %v %T\n", data, data)
	// downcasting
	// c := data.(*City) // ok if you are sure tho have a *City object
	c, ok := data.(*City) // if not ok, c= nil
	fmt.Println("After downcasting:", c, ok)
	if ok {
		fmt.Println(c, c.Name, c.Pop)
	}

	switch v := data.(type) {
	case *City:
		fmt.Println("City:", v, v.Name, v.Pop)
	case *MyUint:
		fmt.Println("MyUint", v, *v, (*v)*2)
	default:
		fmt.Printf("Type inconnu: %v %T", v, v)
	}
}

func testUintIncrementable() {
	var i UintIncrementable
	city := City{"Lyon", 1500000}
	i = &city
	i.Incr(50)
	fmt.Println(city)
	doSomethingWithUintIncrementable(&city)
	fmt.Println(city)
	i = &City{"Toulouse", 1300000}
	i.Incr(1)
	doSomethingWithUintIncrementable(i)

	myuint := MyUint(12)
	i = &myuint
	doSomethingWithUintIncrementable(i)
}

func testStringer() {
	city := City{"Lyon", 1500000}
	// print with String() method from interface Stringer
	fmt.Println(city, &city)
}

func testNew() {
	// new : allocate only
	var city *City = new(City)
	fmt.Println(city, city.Name)

	// &City{..} : allocate and initialize fields
	city = &City{"Lyon", 1500000}
	fmt.Println(city, city.Name)

	// panic(city)

	var intptr = new(int)
	fmt.Println(*intptr)
	*intptr = 12
	*intptr += 2
	(*intptr)++
	fmt.Println(*intptr)
}

func main() {
	// structBasics()
	// addFunctionality()
	// testUintIncrementable()
	// testStringer()
	testNew()
}
