# 对称的二叉树

- 标签: 二叉树, 递归
- 链接: https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
- 难度: 简单
- 题号: 28

### 题目描述

---

请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
```text
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
```

**示例1：**

```
输入：root = [1,2,2,3,4,4,3]
输出：true
```

**示例2：**

```
输入：root = [1,2,2,null,3,null,3]
输出：false
```

### 思路

---

对左右子树同时分别对称遍历。

### 代码

---

**递归法：**

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    } else {
        return recur(root.Left, root.Right)
    }
}

func recur(L, R *TreeNode) bool {
    if L == nil && R == nil {
        return true
    } else if L == nil || R == nil || L.Val != R.Val {
        return false;
    }
    return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}
```

**迭代法：**

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    var leftStack []*TreeNode
    leftStack = append(leftStack, root.Left)
    var rightStack []*TreeNode
    rightStack = append(rightStack, root.Right)

    var leftNode *TreeNode
    var rightNode *TreeNode

    for len(leftStack) != 0 || len(rightStack) != 0 {
        leftNode = leftStack[len(leftStack)-1]
        leftStack = leftStack[:len(leftStack)-1]
        rightNode = rightStack[len(rightStack)-1]
        rightStack = rightStack[:len(rightStack)-1]

        if leftNode == nil && rightNode == nil {
            continue
        } else if leftNode != nil && rightNode != nil {
            leftStack = append(leftStack, leftNode.Left)
            leftStack = append(leftStack, leftNode.Right)
            rightStack = append(rightStack, rightNode.Right)
            rightStack = append(rightStack, rightNode.Left)
        } else {
            return false
        }

        if rightNode.Val != leftNode.Val {
            return false
        }
    }

    return true
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)

- 空间复杂度

    O(n)