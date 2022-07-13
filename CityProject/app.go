package main

import "fmt"

func testNew() {
	// new : allocate only
	var city *City = new(City)
	fmt.Println(city, city.Name)

	// &City{..} : allocate and initialize fields
	city = &City{"Lyon", 1500000}
	fmt.Println(city, city.Name)

}

func main() {
	testNew()
}
