# 调整数组顺序使奇数位于偶数前面

- 标签: 双指针
- 链接: https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
- 难度: 简单
- 题号: 21

### 题目描述

---

输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

**示例1：**

```
输入：nums = [1,2,3,4]
输出：[1,3,2,4] 
注：[3,1,2,4] 也是正确的答案之一。
```

### 思路

---

一前一后指针移动，前指针后移，遇到偶数停下，后指针前移，遇到奇数停下，然后两数交换。

### 代码

---

```go
func exchange(nums []int) []int {
    begin := 0
    end := len(nums) - 1
    for begin < end {
        for nums[begin] % 2 == 1 && begin < end {
            begin++
        }
        for nums[end] % 2 == 0 && begin < end {
            end--
        }
        nums[begin], nums[end] = nums[end], nums[begin]
    }

    return nums
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，完整遍历一遍数组。

- 空间复杂度

    O(1)，额外空间开销只有两个变量。