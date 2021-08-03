package _0_2_numWays

func numWays(n int) int {
	ppre := 1
	pre := 1
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ppre + pre
		if ans >= 1e9+7 {
			ans %= 1e9 + 7
		}
		ppre = pre
		pre = ans
	}

	return ans
}
