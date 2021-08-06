package _6_printNumbers

func printNumbers(n int) []int {
	max := maxNumber(n)
	numbers := make([]int, max-1)
	for i := 1; i < max; i++ {
		numbers[i-1] = i
	}

	return numbers
}

func maxNumber(n int) int {
	product := 1
	x := 10
	for n >= 1 {
		if n%2 == 1 {
			product *= x
			n--
		} else {
			x *= x
			n /= 2
		}
	}

	return product
}
