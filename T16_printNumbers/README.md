# 打印从1到最大的n位数

- 标签: dfs
- 链接: https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
- 难度: 简单
- 题号: 17

### 题目描述

---

输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

**示例1：**

```
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]
```

### 思路

---

此题应该考察大数打印，返回值应该是 []string，但是返回值是 []int 就过于简单，大数打印思路见题解。

此题还需注意一点 append 函数会增加时间和空间开销，注意初始化 slice 后再使用。

### 代码

---

```go
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
        if n % 2 == 1 {
            product *= x
            n--
        } else {
            x *= x
            n /= 2
        }
    }
    
    return product
}
```

### 复杂度分析

---

- 时间复杂度

    O(10^n)。

- 空间复杂度

    O(1)，[]int 为返回值，不算额外开销。