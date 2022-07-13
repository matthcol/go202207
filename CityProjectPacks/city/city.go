package city

import "fmt"

type City struct {
	Name string
	Pop  uint
}

func (city City) ToString() string {
	return fmt.Sprintf("%s (pop: %d)", city.Name, city.Pop)
}

// implementation of interface Stringer
func (city City) String() string {
	return city.ToString()
}

func (city *City) Incr(delta uint) uint {
	city.Pop += delta
	return city.Pop
}
