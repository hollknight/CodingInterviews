package T13_movingCount

func movingCount(m int, n int, k int) int {
	arrays := make([][]bool, m, m)
	count := 0
	for i := 0; i < m; i++ {
		arrays[i] = make([]bool, n, n)
	}
	for i := 0; i < m; i++ {
		arrays[i] = make([]bool, n, n)
		if canMove(i, 0, k) {
			arrays[i][0] = true
			count++
		} else {
			break
		}
	}
	for i := 1; i < n; i++ {
		if canMove(0, i, k) {
			arrays[0][i] = true
			count++
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if (arrays[i-1][j] || arrays[i][j-1]) && canMove(i, j, k) {
				arrays[i][j] = true
				count++
			}
		}
	}

	return count
}

func canMove(x, y, k int) bool {
	xSum := 0
	ySum := 0
	for x > 0 {
		xSum += x % 10
		x /= 10
	}
	for y > 0 {
		ySum += y % 10
		y /= 10
	}

	if k >= xSum+ySum {
		return true
	} else {
		return false
	}
}
