package _7_isMatch

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 分成空正则和非空正则两种
			if j == 0 {
				f[i][j] = (i == 0)
			} else {
				// 非空正则分为两种情况 * 和 非*
				if p[j-1] != '*' {
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						f[i][j] = f[i-1][j-1]
					}
				} else {
					// 碰到 * 了，分为看和不看两种情况
					// 不看
					if j >= 2 {
						f[i][j] = f[i][j] || f[i][j-2]
					}
					// 看
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						f[i][j] = f[i][j] || f[i-1][j]
					}
				}
			}
		}
	}
	return f[n][m]
}
