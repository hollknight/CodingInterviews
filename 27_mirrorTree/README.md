# 二叉树的镜像

- 标签: 递归
- 链接: https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/
- 难度: 简单
- 题号: 27

### 题目描述

---

请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

​    4
  /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：
​     4
   /   \
  7     2
 / \   / \
9   6 3   1

**示例1：**

```
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
```

### 思路

---

遍历二叉树时交换节点即可。

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
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    mirrorRoot := new(TreeNode)
    mirrorRoot.Val = root.Val
    mirrorRoot.Left = mirrorTree(root.Right)
    mirrorRoot.Right = mirrorTree(root.Left)

    return mirrorRoot
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)

- 空间复杂度

    O(n)，当二叉树退化为链表时递归深度为n。