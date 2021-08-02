package _4_1_cuttingRope

import "math"

func cuttingRope(n int) int {
	max := 1
	var product int
	for i := 2; i <= n; i++ {
		num := n / i
		rest := n % i
		if num == 1 {
			product = i * rest
		} else {
			temp1 := int(math.Pow(float64(i), float64(num-1))) * (i + rest)
			temp2 := int(math.Pow(float64(i), float64(num))) * rest
			if temp1 > temp2 {
				product = temp1
			} else {
				product = temp2
			}
		}
		if product > max {
			max = product
		}
	}

	return max
}
