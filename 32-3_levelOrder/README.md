# 从上到下打印二叉树 III

- 标签: 动态规划
- 链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
- 难度: 中等
- 题号: 32 - III

### 题目描述

---

请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

**示例1：**

```
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]
```

### 思路

---

记录奇偶数层数，偶数层层序遍历时反转切片。

### 代码

---

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var queue []*TreeNode
    queue = append(queue, root)
    var ans [][]int
    rowFlag := 0

    for len(queue) != 0 {
        var row []int
        for i := len(queue); i > 0; i-- {
            var node *TreeNode
            node = queue[0]
            queue = queue[1:]
            
            if node != nil {
                queue = append(queue, node.Left)
                queue = append(queue, node.Right)
                row = append(row, node.Val)
            }
        }
        if rowFlag % 2 == 1 {
            row = reverse(row)
        }
        rowFlag++
        ans = append(ans, row)
    }

    return ans[:len(ans)-1]
}

func reverse(slice []int) []int {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
    return slice
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)

- 空间复杂度

    O(n)