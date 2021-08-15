# 二叉搜索树的后序遍历序列

- 标签: 栈
- 链接: https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/
- 难度: 中等
- 题号: 33

### 题目描述

---

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

```
参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
```

**示例1：**

```
输入: [1,3,2,6,5]
输出: true
```

**示例2：**

```
输入: [1,6,3,2,5]
输出: false
```

### 思路

---

见k神题解。

[力扣题解](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/solution/mian-shi-ti-33-er-cha-sou-suo-shu-de-hou-xu-bian-6/)

### 代码

---

```go
func verifyPostorder(postorder []int) bool {
    var stack []int
    root := math.MaxInt32
    for i := len(postorder)-1; i >= 0; i-- {
        if (postorder[i] > root) {
            return false
        }
        for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, postorder[i])
    }

    return true
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，各节点入栈一次出栈一次。

- 空间复杂度

    O(n)。