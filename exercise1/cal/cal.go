package cal

func isEven(n int) bool {
	return n % 2 == 0
}

func SumOfEven(n int) int {
	var sum int = 0
	for i := 0; i <= n; i++ {
		if isEven(i) {
			sum += i
		} else {
			sum += 0
		}
	}
	return sum
}