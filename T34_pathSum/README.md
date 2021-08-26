# 二叉树中和为某一值的路径

- 标签: dfs, 二叉树
- 链接: https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
- 难度: 中等
- 题号: 34

### 题目描述

---

输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

**示例1：**

```
给定如下二叉树，以及目标和 target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
```

### 思路

---

dfs前序遍历一次，注意中间结果的储存。

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
func pathSum(root *TreeNode, target int) [][]int {
    if root == nil {
        return nil
    }
    var nums []int
    var ans [][]int
    findSum(root, target, nums, &ans)
    
    return ans
}

func findSum(root *TreeNode, target int, nums []int, ans *[][]int) {
    if root == nil {
        return
    }

    nums = append(nums, root.Val)
    target -= root.Val
    if target == 0 && root.Left == nil && root.Right == nil {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        *ans = append(*ans, tmp)
    }
    findSum(root.Left, target, nums, ans)
    findSum(root.Right, target, nums, ans)
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(n)。