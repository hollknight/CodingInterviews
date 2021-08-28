# 连续子数组的最大和

- 标签: 动态规划
- 链接: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
- 难度: 简单
- 题号: 42

### 题目描述

---

输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

**示例1：**

```
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

- `1 <= arr.length <= 10^5`
- `100 <= arr[i] <= 100`

### 思路

---

使用temp变量按顺序遍历，当temp为正数就继续叠甲，当temp是负数就置0。在遍历过程中，如果temp > max，就更新max。但是在更新的过程中，要注意最大子序列和为负数和0的情况，在这种情况下不能简单地更新max = 0,。

### 代码

---

```go
func maxSubArray(nums []int) int {
    max := nums[0]
    temp := 0
    for i := 0; i < len(nums); i++ {
        temp += nums[i]
        if temp > max {
            if temp == 0 {
                if nums[i] == 0 {
                    max = temp
                }
                continue
            }
            max = temp
        }
                if temp < 0 {
            temp = 0
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