package main

import (
	"fmt"
	"math"
)

// const declaration with inference
const Pi = 3.1459

// const declaration with type
// const Pi float32 = 3.1459

func basicTypes() {
	var temperature int8
	var distance uint64
	var city2 string = "Pau"
	var speed float64 = 49.99
	price := 98
	city := "Lyon"
	temperature = 22
	distance = 5000000000
	fmt.Println("It's monday")
	fmt.Println(temperature)
	fmt.Println(distance)
	fmt.Println(price)
	fmt.Println(city)
	fmt.Println(city2)
	fmt.Println(speed)
	speed = 0.1 // 0.000110011001100110011.. (base 2)
	fmt.Printf("Speeds: %f, %f, %f\n", speed, 2*speed, 3*speed)
	fmt.Printf("Speeds: %.18f, %.18f, %.18f\n", speed, 2*speed, 3*speed)
	speed = 1e308
	fmt.Printf("Max speed: %f, infinite speed: %f\n", speed, 2*speed)

	var area float32 = Pi * float32(math.Pow(12.2, 2))
	fmt.Println(area)
	var perimeter float64 = 2.0 * Pi * 12.2
	fmt.Println(perimeter)
}

func temperatureMin() int8 {
	return 20
}

func temperatureMinMax() (int8, int8) {
	return 20, 44
}

func calcul(t1, t2 int8) int8 {
	return t2*2 - t1
}

func temperatureDay() (int8, string) {
	return 25, "Monday"
}

func testFunctions() {
	temp := temperatureMin()
	tMin, tMax := temperatureMinMax()
	res := calcul(tMin, tMax)
	fmt.Println(temp)
	fmt.Printf("Min: %d, Max: %d\n", tMin, tMax)
	fmt.Printf("Result: %d\n", res)
	tempDay, day := temperatureDay()
	fmt.Printf("Temperature is %d %s\n", tempDay, day)
	var tempDay2 int8
	var day2 string
	tempDay2, day2 = temperatureDay()
	fmt.Printf("Temperature is %d %s\n", tempDay2, day2)
}

func displaySlice(slice []int8) {
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		slice,
		len(slice),
		cap(slice))
}

func constructArraySlice() {
	// array : constant size
	var temperatures [4]int8
	temperatures[0] = 22
	temperatures[1] = 25
	// temperatures[2] = 42
	temperatures[3] = 38
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperatures,
		len(temperatures),
		cap(temperatures))

	// slice : array of dynamic size
	// complex declaration
	var temperatures2019 []int8 = []int8{23, 26, 42, 38}
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperatures2019,
		len(temperatures2019),
		cap(temperatures2019))

	// with inference
	var temperatures2020 = []int8{24, 21, 44, 26}
	temperatures2021 := []int8{24, 21, 44, 26}
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperatures2020,
		len(temperatures2020),
		cap(temperatures2020))
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperatures2021,
		len(temperatures2021),
		cap(temperatures2021))

	// append: copy original data on the new slice
	temperaturesNew := append(temperatures2019, 28, -1, 32)
	temperatures2019[0]++
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperatures2019,
		len(temperatures2019),
		cap(temperatures2019))
	fmt.Printf("values: %v, length: %d,  capacity: %d\n",
		temperaturesNew,
		len(temperaturesNew),
		cap(temperaturesNew))

	// a slice shares its memory with original data
	extrait := temperaturesNew[2:5]
	temperaturesNew[2]++
	displaySlice(extrait)

	temperaturesNew = temperaturesNew[5:]
	displaySlice(temperaturesNew)

	// slice on a whole array
	extrait = temperatures[:]
	temperatures[0]++
	displaySlice(extrait)

	// first 2 values
	extrait = temperatures2020[:2]
	displaySlice(extrait)

	// bound out of range
	// temperatures2020 = temperatures2020[2:7]
	// displaySlice(temperatures2020)

	// OK because length was 2 and capacity 3
	extrait = temperaturesNew[:3]
	displaySlice(extrait)

	aLotOfTemperatures := make([]int8, 5)
	displaySlice(aLotOfTemperatures)

	// create a slice of length 0 with 1M capacity
	aLotOfTemperatures2 := make([]int8, 0, 1000000)
	displaySlice(aLotOfTemperatures2)
	fmt.Println(aLotOfTemperatures2 == nil)

	// extend slice by 1
	aLotOfTemperatures2 = aLotOfTemperatures2[:1]
	aLotOfTemperatures2[0] = 23
	displaySlice(aLotOfTemperatures2)

	// extend slice by 3
	aLotOfTemperatures2 = aLotOfTemperatures2[:len(aLotOfTemperatures2)+3]
	aLotOfTemperatures2[3] = 44
	displaySlice(aLotOfTemperatures2)

	// reduce slice
	aLotOfTemperatures2 = aLotOfTemperatures2[:3]
	displaySlice(aLotOfTemperatures2)

	// reduce slice
	aLotOfTemperatures2 = aLotOfTemperatures2[1:]
	displaySlice(aLotOfTemperatures2)

	// nil slice
	// nilTemperature := make([]int8, 0, 0) // is not nil
	var nilTemperature []int8 // a true nil
	fmt.Println((nilTemperature == nil))

	extrait = nil
	fmt.Println((extrait == nil))
	displaySlice(extrait)

	// to extend nil slice use make, append
	extrait = append(extrait, 22, 32)
	// extrait = make([]int8, 4, 10000)
	displaySlice(extrait)
}

func loopsWithSlices() {
	temperatures := make([]int8, 20)
	for i := range temperatures {
		temperatures[i] = int8(i)*2 + 3
	}
	displaySlice(temperatures)
	// loop on indexes + values
	for i, v := range temperatures {
		fmt.Printf(" - t[%d] = %d\n", i, v)
	}

	for i := range temperatures {
		temperatures[i] += 2
	}

	// loop on values only
	for _, v := range temperatures {
		fmt.Printf(" * %d\n", v)
	}
}

func playWithStrings() {
	city := "Hà Nội"
	city2 := "福島市"
	fmt.Println(city, city2)
	// len(string) => nb of bytes
	fmt.Println(len(city), len(city2))
	cityRunes := []rune(city)
	cityRunes2 := []rune(city2)
	fmt.Println(len(cityRunes), len(cityRunes2))
	fmt.Println(cityRunes)
	fmt.Println(cityRunes2)
	for _, c := range cityRunes2 {
		fmt.Printf("# %d : %c ", c, c)
	}
}

func loops() {
	// see other examples

	for i := 0; i < 10; i++ {
		fmt.Println("-", i)
	}
	// fmt.Println(i) // i is not accessible after loop

	for i := 0; i < 10; {
		fmt.Println("*", i)
		i++
	}

	var j = 0
	for j < 10 {
		fmt.Println("~", j)
		if j == 5 {
			break
		}
		j++
	}
	fmt.Println("j after loop", j)

}

func isEmpty(slice []int) bool {
	return len(slice) == 0
}

func ifs() {
	temperatures := []int{12, 34}
	if b := isEmpty(temperatures); b {
		fmt.Println("empty slice", b)
	} else {
		fmt.Println("slice with temperatures", b)
	}

	if l := len(temperatures); l > 0 {
		fmt.Println("slice with temperatures", l)
	} else {
		fmt.Println("empty slice", l)
	}
	// fmt.Println(l) // l not defined after if
}

func cases(n int) {
	switch v := 5; v {
	case 1:
		fmt.Println("case 1")
	case 2 + n:
		fmt.Println("case 2 + n")
	case 5:
		fmt.Println("case 5")
	case 8:
		fmt.Println("case 8")
	default:
		fmt.Println("case default")
	}
}

func maps() {
	// nil map
	// var popByCity map[string]uint

	// map initialized
	var popByCity = make(map[string]uint)
	popByCity["Lyon"] = 1700000
	fmt.Println(popByCity)

	var popByCity2 = map[string]uint{
		"Lyon":     1700000,
		"Toulouse": 1300000,
	}
	popByCity2["Clermont-Ferrand"] = 300000
	fmt.Println(popByCity2)

	popByCity2["Clermont-Ferrand"]++
	delete(popByCity2, "Toulouse")
	fmt.Println(popByCity2)

	pop := popByCity2["Clermont-Ferrand"]
	fmt.Println("Pop Clermont-Ferrand:", pop)

	pop = popByCity2["Toulouse"]
	fmt.Println("Pop Toulouse:", pop)

	pop2, ok := popByCity2["Toulouse"]
	fmt.Println("Pop Toulouse:", pop2, ok)
	fmt.Printf("Pop Toulouse: %d (%t)\n", pop2, ok)

	for c, p := range popByCity2 {
		fmt.Printf(" - %s : %d\n", c, p)
	}

	for c := range popByCity2 {
		fmt.Printf(" * %s\n", c)
	}

	for _, p := range popByCity2 {
		fmt.Printf(" ~ %d\n", p)
	}

}

func main() {
	// basicTypes()
	// testFunctions()
	// constructArraySlice()
	// loopsWithSlices()
	// playWithStrings()
	// loops()
	// ifs()
	// cases(0)
	// cases(3)
	maps()
}
