# 机器人的运动范围

- 标签: 动态规划
- 链接: https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
- 难度: 中等
- 题号: 13

### 题目描述

---

地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

**示例1：**

```text
输入：m = 2, n = 3, k = 1
输出：3
```

- 提示

    1 <= n,m <= 100

    0 <= k <= 20

### 思路

---

动态规划记录机器人已经走过的格子，此题还有个隐藏条件：探测时只需要向右和下两个方向探测即可。

### 代码

---

```go
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

    if k >= xSum + ySum {
        return true
    } else {
        return false
    }
}
```

### 复杂度分析

---

- 时间复杂度

    O(n^2)。

- 空间复杂度

    O(n^2)