package backup

import (
	"cityapp/city"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestParseCities(t *testing.T) {
	textCities := `Lyon (pop: 1500000)
	Toulouse (pop: 1300000)
	Pau (pop: 100000)
	`
	expectedCities := []city.City{
		{Name: "Lyon", Pop: 1500000},
		{Name: "Toulouse", Pop: 1300000},
		{Name: "Pau", Pop: 100000},
	}

	actualCities, err := ParseCities(textCities)
	if err != nil {
		t.Errorf("ParseCities not error was expected, got %s", err)
	}
	for i, actualCity := range actualCities {
		expectedCity := expectedCities[i]
		if actualCity.Name != expectedCity.Name {
			t.Errorf("ParseCities got city name %s, got %s",
				actualCity.Name, expectedCity.Name)
		}
	}

}

func TestReaderWriter(t *testing.T) {
	textCities := `Lyon (pop: 1500001)
	Toulouse (pop: 1300001)
	Pau (pop: 100001)
	`
	reader := strings.NewReader(textCities)
	writer, err := os.Create("cities.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()
	io.Copy(writer, reader)
	t.Log("copy done")
}
