# 顺时针打印矩阵

标签: 数组
链接: https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
难度: 简单
题号: 29

### 题目描述

---

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

**示例1：**

```
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
```

### 思路

---

记录每一行、每一列的遍历起始于终止位置，每一次遍历更新起止位置。

### 代码

---

```go
func spiralOrder(matrix [][]int) []int {
    var nums []int
    rowBegin := 0
    rowEnd := len(matrix) - 1
    if rowEnd == -1 {
        return nums
    }
    columBegin := 0
    columEnd := len(matrix[0]) - 1

    for columBegin <= columEnd && rowBegin <= rowEnd {
        for i := columBegin; i <= columEnd; i++ {
            nums = append(nums, matrix[rowBegin][i]) 
        }
        rowBegin++
        for i := rowBegin; i <= rowEnd; i++ {
            nums = append(nums, matrix[i][columEnd])
        }
        columEnd--
        for i := columEnd; i >= columBegin; i-- {
            nums = append(nums, matrix[rowEnd][i])
        }
        rowEnd--
        for i := rowEnd; i >=rowBegin; i-- {
            nums = append(nums, matrix[i][columBegin])
        }
        columBegin++
    }
    nums = nums[:len(matrix)*len(matrix[0])]
     
     return nums
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(1)。