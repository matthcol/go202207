package city

import (
	"encoding/json"
	"testing"
)

func TestToString(t *testing.T) {
	name := "Lyon"
	var pop uint = 1500000
	city1 := City{Name: name, Pop: pop}
	actualText := city1.ToString()
	expectedText := "Lyon (pop: 1500000)"
	if actualText != expectedText {
		// t.Fail()
		t.Errorf("ToString incorrect, got %s, want %s",
			actualText, expectedText)
	}

}

func TestJsonMarshalling(t *testing.T) {
	city := City{Name: "Lyon", Pop: 1500000}
	bjson, _ := json.Marshal(city)
	t.Log(string(bjson))
}
