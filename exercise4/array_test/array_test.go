package arraytest

import "testing"

func TestMaxDiff_given5_expected998(t *testing.T){
	// Arrange
	num := []int{1, 999, 2, 3, 4}
	expected := 998

	// Act
	actual := Difference(num)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}
}

func TestIsTwo_given2_expected997(t *testing.T){
	// Arrange
	num := []int{999, 2}
	expected := 997

	// Act
	actual := Difference(num)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}
}

func TestLessThenTwo_given1_expected0(t *testing.T){
	// Arrange
	num := []int{1}
	expected := 0

	// Act
	actual := Difference(num)

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}
}