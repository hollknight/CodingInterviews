# 剪绳子

- 标签: 贪心
- 链接: https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
- 难度: 中等
- 题号: 14- I

### 题目描述

---

给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

**示例1：**

```text
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
```

- 提示

    2 ≤ n ≤ 58

### 思路

---

用贪心的思想，数字分的较为平均则乘积较大。

更为严谨正确的数学证明见题解：

[力扣](https://leetcode-cn.com/problems/jian-sheng-zi-lcof/solution/mian-shi-ti-14-i-jian-sheng-zi-tan-xin-si-xiang-by/)

### 代码

---

```go
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
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(1)。