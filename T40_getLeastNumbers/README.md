# 最小的k个数

- 标签: 递归
- 链接: https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
- 难度: 简单
- 题号: 40

### 题目描述

---

输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

**示例1：**

```
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
```

### 思路

---

因为题目只要求找到最小的 k 个数，不要求有序。因此，我们可以在快速排序的基础上更改，让哨兵只整理需要重新整理的一边。

[力扣](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/jian-zhi-offer-40-zui-xiao-de-k-ge-shu-j-9yze/)

### 代码

---

```go
func getLeastNumbers(arr []int, k int) []int {
    quickSort(&arr, 0, len(arr)-1, k)
    return arr[:k]
}

func quickSort(arr *[]int, left, right, k int) {
    if left >= right {
        return
    }
    i := left
    j := right
    for i < j {
        for i < j && (*arr)[j] >= (*arr)[left] {
            j--
        }
        for i < j && (*arr)[i] <= (*arr)[left] {
            i++
        }
        (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
    }
    (*arr)[left], (*arr)[i] = (*arr)[i], (*arr)[left]
    if i > k {
        quickSort(arr, left, i-1, k)
    } else if i < k {
        quickSort(arr, i+1, right, k)
    }
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，具体分析见题解。

- 空间复杂度

    O(logn)，平均递归深度为logn。