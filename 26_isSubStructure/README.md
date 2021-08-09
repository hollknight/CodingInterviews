# 树的子结构

- 标签: 二叉树, 递归
- 链接: https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
- 难度: 中等
- 题号: 26

### 题目描述

---

输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:
     3
    / \
   4   5
  / \
 1   2

给定的树 B：
   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

**示例1：**

```
输入：A = [3,4,5,1,2], B = [4,1]
输出：true
```

### 思路

---

1. 如果传入的 B 节点是根节点
    - A.Val == B.Val

        继续往后判断 || 判断 A 的左子树或右子树中是否含有 B

    - A.Val != B.Val

        判断 A 的左子树或右子树中是否含有 B

2. 如果传入的 B 节点不是根节点
    - A.Val == B.Val

        继续往后判断

    - A.Val != B.Val

        return false

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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    return is(A, B, true)
}

func is(A *TreeNode, B *TreeNode, isRoot bool) bool {
    if A == nil || B == nil {
        if A == B {
            return true
        } else {
            return false
        }
    }

    if A.Val == B.Val {
        if B.Left == nil {
            return is(A.Right, B.Right, false) || is(A.Left, B, true) || is(A.Right, B, true)
        } else if B.Right == nil {
            return is(A.Left, B.Left, false) || is(A.Left, B, true) || is(A.Right, B, true)
        } else {
            return is(A.Left, B.Left, false) && is(A.Right, B.Right, false) || is(A.Left, B, true) || is(A.Right, B, true)
        }
    } else {
        if isRoot {
            return is(A.Left, B, true) || is(A.Right, B, true)
        } else {
            return false
        }
    }
}
```

### 复杂度分析

---

- 时间复杂度

    O(mn)， 其中 m,n 分别为树 A 和 树 B 的节点数量；先序遍历树 A 占用 O(m) ，每次判断子树中是否含有子结构占用 O(N) 。

- 空间复杂度

    O(m)，

    当树 A 和树 B 都退化为链表时，递归调用深度最大。m≤n 时，遍历树 A 与递归判断的总递归深度为 m ；当 m>n 时，最差情况为遍历至树 A 叶子节点，此时总递归深度为 m。