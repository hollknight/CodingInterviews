# 重建二叉树

- 标签: 递归
- 链接: https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
- 难度: 中等
- 题号: 07

### 题目描述

---

输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。

假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

**示例1：**

![https://assets.leetcode.com/uploads/2021/02/19/tree.jpg](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

### 思路

---

根据前序遍历的数组结果获取根节点值，根据根节点值在中序遍历中获取左子树和右子树节点数目，之后分别根据长度截取左子树和右子树的前序遍历与中序遍历结果。递归建成整个树。

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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    root := new(TreeNode)
    root.Val = preorder[0]

    index := 0
    for ; index < len(inorder); index++ {
        if inorder[index] == preorder[0] {
            break
        }
    }

    if index > 0 {
        root.Left = buildTree(preorder[1:index+1], inorder[:index])
    } else {
        root.Left = buildTree(nil, nil)
    }
    
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])

    return root
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，n 为节点数目。

- 空间复杂度

    O(n)，递归消耗的栈空间。