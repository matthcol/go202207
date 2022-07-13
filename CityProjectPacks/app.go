package main

import (
	"cityapp/city"
	"cityapp/persistence"
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

func main() {
	testNew()
	persistence.TestDrivers()
}
