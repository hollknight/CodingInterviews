# 从上到下打印二叉树

- 标签: 二叉树
- 链接: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
- 难度: 中等
- 题号: 32 - I

### 题目描述

---

从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

**示例1：**

```
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回：
[3,9,20,15,7]
```

### 思路

---

使用队列 FIFO 的特性，将节点依次放入队列并拿出。

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
func levelOrder(root *TreeNode) []int {
    var ans []int
    var queue []*TreeNode
    queue = append(queue, root)
    
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        if node != nil {
            ans = append(ans, node.Val)
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
    }

    return ans
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(n)。