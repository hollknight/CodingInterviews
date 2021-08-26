package T14_2_cuttingRope

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	b := n % 3
	var p uint64 = 1e9 + 7
	var rem uint64 = 1
	var x uint64 = 3
	for i := n/3 - 1; i > 0; i /= 2 {
		if i%2 == 1 {
			rem = (rem * x) % p
		}
		x = (x * x) % p
	}
	if b == 0 {
		return int(rem * 3 % p)
	}
	if b == 1 {
		return int(rem * 4 % p)
	}

	return int(rem * 6 % p)
}
