# 二叉树的深度

- 标签: dfs, 二叉树
- 链接: https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
- 难度: 简单
- 题号: 55 - I

### 题目描述

---

输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

**示例1：**

```go
例如：

给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
```

### 思路

---

后序遍历一次完整的树，遍历时记录更新最大深度。

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
func maxDepth(root *TreeNode) int {
    maxDepth := 0
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, depth int) {
        if root == nil {
            return
        }
        if depth > maxDepth {
            maxDepth = depth
        }
        dfs(root.Left, depth+1)
        dfs(root.Right, depth+1)
    }
    dfs(root, 1)
    return maxDepth
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(n)。