package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func sortInts() {
	data := []int{12, 3, 5, 25}
	sort.Ints(data)
	fmt.Println(data)
}

func sortFloats() {
	data := []float64{12.2, 3e10, -5e-2,
		math.Inf(1), math.Inf(-1), math.NaN()}
	sort.Float64s(data)
	fmt.Println(data)
}

func sortStrings() {
	data := []string{"Lyon", "Toulouse", "Clermont-Ferrand", "Pau"}
	sort.Strings(data)
	fmt.Println(data)
	fmt.Println(strings.Compare("Lyon", "Toulouse"))
}

type SortableIgnoreCase []string

func (sortable SortableIgnoreCase) Len() int {
	return len(sortable)
}

func (sortable SortableIgnoreCase) Less(i, j int) bool {
	return strings.Compare(
		strings.ToLower(sortable[i]),
		strings.ToLower(sortable[j])) < 0
}

func (sortable SortableIgnoreCase) Swap(i, j int) {
	sortable[i], sortable[j] = sortable[j], sortable[i]
}

type SortableFrench []string

func (sortable SortableFrench) Len() int {
	return len(sortable)
}

func (sortable SortableFrench) Less(i, j int) bool {
	collator := collate.New(language.French)
	return collator.CompareString(
		sortable[i],
		sortable[j]) < 0
}

func (sortable SortableFrench) Swap(i, j int) {
	sortable[i], sortable[j] = sortable[j], sortable[i]
}

func sortStringsIgnoreCase() {
	data := SortableIgnoreCase{"Lyon", "Toulouse", "Clermont-Ferrand", "pau"}
	sort.Sort(data)
	fmt.Println(data)
}

func sortStringsFrenchAccent() {
	data := SortableFrench{"été", "étage", "étuve"}
	sort.Sort(data)
	fmt.Println(data)
}

func sortStringsSpanish() {
	data := []string{"mañana", "mano", "matador"}
	collatorEs := collate.New(language.Spanish)
	lessEs := func(i, j int) bool {
		return collatorEs.CompareString(data[i], data[j]) < 0
	}
	sort.Slice(data, lessEs)
	fmt.Println(data)
}

func main() {
	// sortInts()
	// sortFloats()
	// sortStrings()
	sortStringsIgnoreCase()
	sortStringsFrenchAccent()
	sortStringsSpanish()
}
