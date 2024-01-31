package cal

import "testing"

func TestSumOfEven_4_expected_6(t *testing.T) {
	// AAA Patern - Arrange, Act, Assert
	
	// Arrange
	given := 4
	expected := 6

	// Act
	actual := SumOfEven(given)

	// Assert
	if actual != expected {
		t.Errorf("SumOfEven(%d) = %d; expected %d", given, actual, expected)
	}
}