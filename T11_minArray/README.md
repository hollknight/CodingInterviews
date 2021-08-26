# 旋转数组的最小数字


- 标签: 二分法
- 链接: https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
- 难度: 简单
- 题号: 11

### 题目描述

---

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

**示例1：**

```
输入：[3,4,5,1,2]
输出：1
```

**示例2：**

```
输入：[2,2,2,0,1]
输出：0
```

### 思路

---

二分法：

![https://pic.leetcode-cn.com/1599404042-JMvjtL-Picture1.png](https://pic.leetcode-cn.com/1599404042-JMvjtL-Picture1.png)

过程见代码注释。

### 代码

---

```go
func minArray(numbers []int) int {
    begin := 0
    end := len(numbers) - 1
    mid := (begin + end) / 2
    
    for begin <= end {
				// 如果最左端元素小于最右端元素，则此时左到右递增，最小元素即为最左端元素
        if numbers[begin] < numbers[end] {
            return numbers[begin]
        }
				// 如果中心元素大于最左端元素，则中心元素一定在最小值左侧，将最左端移到mid+1处
        if numbers[mid] > numbers[begin] {
            begin = mid + 1
        }
				// 如果中心元素小于最左端元素，则中心元素在最小值处或最小值右侧，将最右端移到mid处
				else if numbers[mid] < numbers[begin] {
            end = mid
        }
				// 此时中心元素等于最左侧元素，将最左端右移
				else {
            begin++
        }
        mid = (begin + end) / 2
    }
    mid = (begin + end) / 2

    return numbers[mid]
}
```

### 复杂度分析

---

- 时间复杂度

    O(logn)。

- 空间复杂度

    O(1)。