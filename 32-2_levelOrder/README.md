# 从上到下打印二叉树 II

- 标签: 动态规划
- 链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
- 难度: 简单
- 题号: 32 - II

### 题目描述

---

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

**示例1：**

```text
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
  [9,20],
  [15,7]
]
```

### 思路

---

层序遍历的同时，每层单独循环，然后储存。

有个小技巧，len(queue) 在循环过程中会发生变化，因此可以从高往低遍历。

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

    for len(queue) != 0 {
        var row []int
        for i := len(queue); i > 0; i-- {
            node := queue[0]
            queue = queue[1:]
            if node != nil {
                queue = append(queue, node.Left)
                queue = append(queue, node.Right)
                row = append(row, node.Val)
            }
        }
        ans = append(ans, row)
    }

    return ans[:len(ans)-1]
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(n)。