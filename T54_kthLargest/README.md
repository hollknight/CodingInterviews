# 二叉搜索树的第k大节点

- 标签: dfs, 二叉树
- 链接: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/
- 难度: 简单
- 题号: 54

### 题目描述

---

给定一棵二叉搜索树，请找出其中第k大的节点。

**示例1：**

```go
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
```

### 思路

---

根据二叉搜索树的性质，我们使用右根左的遍历方式，每次遍历k—，当k==0时即为第k大数，在返回条件处加上k≤0可以避免多余的遍历。

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
func kthLargest(root *TreeNode, k int) int {
    ans := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil || k <= 0 {
            return
        }
        dfs(root.Right)
        k--
        if k == 0 {
            ans = root.Val
            return
        }
        dfs(root.Left)
    }
    dfs(root)
    return ans
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，二叉树退化为链表。

- 空间复杂度

    O(n)，二叉树退化为链表。