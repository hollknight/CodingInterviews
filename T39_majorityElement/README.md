# 数组中出现次数超过一半的数字

- 标签: 数组
- 链接: https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
- 难度: 简单
- 题号: 39

### 题目描述

---

数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

**示例1：**

```
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2
```

### 思路

---

[力扣](https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/solution/mian-shi-ti-39-shu-zu-zhong-chu-xian-ci-shu-chao-3/)

### 代码

---

```go
func majorityElement(nums []int) int {
    vote := 0
    majority := 0
    for i := 0; i < len(nums); i++ {
        if vote == 0 {
            majority = nums[i]
        }
        if majority == nums[i] {
            vote++
        } else {
            vote--
        }
    }

    return majority
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(1)。