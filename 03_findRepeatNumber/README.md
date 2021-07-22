# 数组中重复的数字

- 备注: 原地交换法
- 来源: 剑指 Offer
- 标签: 数组
- 链接: https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
- 难度: 简单
- 题号: 03

### 题目描述

---

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

**示例1：**

```text
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
```

### 思路

---

- 哈希表

- 原地交换法

  因为题目中规定数组长度为`n`，所有数字都在`0 ~ n-1`范围内，因此我们可以将整个数组看作为一个哈希表。

  ![](https://pic.leetcode-cn.com/1618146573-bOieFQ-Picture0.png)

  遍历中，第一次遇到数字 xx 时，将其交换至索引 xx 处；而当第二次遇到数字 xx 时，一定有 nums[x] = xnums[x]=x ，此时即可得到一组重复数字。

  算法流程：
  遍历数组 numsnums ，设索引初始值为 i = 0i=0 :

  若 nums[i] = inums[i]=i ： 说明此数字已在对应索引位置，无需交换，因此跳过；
  若 nums[nums[i]] = nums[i]nums[nums[i]]=nums[i] ： 代表索引 nums[i]nums[i] 处和索引 ii 处的元素值都为 nums[i]nums[i] ，即找到一组重复值，返回此值 nums[i]nums[i] ；
  否则： 交换索引为 ii 和 nums[i]nums[i] 的元素值，将此数字交换至对应索引位置。
  若遍历完毕尚未返回，则返回 -1−1 。

### 代码

---

- 哈希表法：

```go
func findRepeatNumber(nums []int) int {
    numMap := make(map[int]bool)
    for _, num := range nums {
        if _, ok := numMap[num]; ok {
            return num
        }
        numMap[num] = true
    }

    return 0
}
```

- 原地交换法：

```go
func findRepeatNumber(nums []int) int {
    i := 0
    for i < len(nums) {
        if nums[i] == i {
            i++
            continue
        }
        if nums[nums[i]] == nums[i] {
            return nums[i]
        } 
        tmp := nums[i]
        nums[i] = nums[tmp]
        nums[tmp] = tmp
    }
    return -1
}
```

### 复杂度分析

---

- 时间复杂度

  两种方法均为O(n)。

- 空间复杂度

  哈希表法为O(n)，原地交换法为O(1)。