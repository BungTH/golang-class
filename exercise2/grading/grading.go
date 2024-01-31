package grading

func Grade(score int) string {
	if score > 100 {
		return "Invalid score"
	} else if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else if score >= 0{
		return "F"
	}
	return "Invalid score"
}