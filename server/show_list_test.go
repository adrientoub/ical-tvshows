package server

import (
	"fmt"
	"testing"
)

func testParse(t *testing.T, expected string, test string) {
	if title := ParseTitle(test); title != expected {
		t.Errorf("Should be `%s`, but is `%s`", expected, title)
	}
}

func TestParseTitle(t *testing.T) {
	testParse(t, "The Flash", "The Flash (2014)")
	testParse(t, "The Flash", "The Flash")
	testParse(t, "Bull", "Bull (2016)")
	testParse(t, "Bull", "Bull")
}

func ExampleParseTitle() {
	fmt.Println(ParseTitle("The Flash (2016)"))
	fmt.Println(ParseTitle("The Flash"))
	// Output:
	// The Flash
	// The Flash
}
