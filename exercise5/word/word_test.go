package word

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	// Arrange
	given := "word word hello hi word"
	expected := map[string]int{
		"word": 3,
		"hello": 1,
		"hi": 1,
	}

	// Act
	actual := WordCount(given)

	// Assert
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("WordCount(%q) = %v; expected %v", given, actual, expected)
	}
}