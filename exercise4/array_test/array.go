package arraytest

func Difference(num []int) int {
	min := num[0]
	max := num[0]
	for i := 0; i < len(num); i++ {
		if num[i] < min {
			min = num[i]
		}
		if num[i] > max {
			max = num[i]
		}
	}
	return max - min
}