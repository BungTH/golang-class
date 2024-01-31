package grading

import "testing"

func TestGrading_101_expected_invalid(t *testing.T) {
	//Arrage
	given := 101
	expected := "Invalid score"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}

func TestGrading_80_expected_A(t *testing.T) {
	//Arrage
	given := 80
	expected := "A"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}	

func TestGrading_70_expected_B(t *testing.T) {
	//Arrage
	given := 70
	expected := "B"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}

func TestGrading_60_expected_C(t *testing.T) {
	//Arrage
	given := 60
	expected := "C"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}

func TestGrading_50_expected_D(t *testing.T) {
	//Arrage
	given := 50
	expected := "D"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}

func TestGrading_49_expected_F(t *testing.T) {
	//Arrage
	given := 49
	expected := "F"

	//Act
	actual := Grade(given)

	//Assert	
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}

func TestGrading_Negative_expected_invalid(t *testing.T) {
	//Arrage
	given := -1
	expected := "Invalid score"

	//Act
	actual := Grade(given)

	//Assert
	if actual != expected {
		t.Errorf("Grade(%d) = %s; expected %s", given, actual, expected)
	}
}