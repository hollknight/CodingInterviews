package _0_1_fib

func fib(n int) int {
	if n == 0 {
		return 0
	}

	pre := 0
	ppre := 1
	ans := 1

	for i := 1; i <= n; i++ {
		ans = ppre + pre
		ans %= 1e9 + 7
		ppre = pre
		pre = ans
	}

	return ans
}
