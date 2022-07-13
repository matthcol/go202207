package main

import (
	"cityapp/backup"
	"cityapp/city"
	"fmt"
)

func testNew() {
	// new : allocate only
	var city1 *city.City = new(city.City)
	fmt.Println(city1, city1.Name)

	// &City{..} : allocate and initialize fields
	city1 = &city.City{"Lyon", 1500000}
	fmt.Println(city1, city1.Name)

}

func f(a ...int) {
	fmt.Println("All args", a)
	for _, v := range a {
		fmt.Println("-", v)
	}
}

func main() {
	// testNew()
	// persistence.TestDrivers()
	// persistence.TestConnection()
	// f(1, 2, 3)
	// backup.TestPaths()
	// backup.TestOS()
	backup.TestReadWriteCities()
}
