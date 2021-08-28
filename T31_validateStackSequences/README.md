# 栈的压入、弹出序列

- 标签: 模拟
- 链接: https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
- 难度: 中等
- 题号: 31

### 题目描述

---

输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

**示例1：**

```
输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
输出：true
解释：我们可以按以下顺序执行：
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
```

**示例2：**

```
输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
输出：false
解释：1 不能在 2 之前弹出。
```

### 思路

---

考虑借用一个辅助栈 stackstack ，模拟 压入 / 弹出操作的排列。根据是否模拟成功，即可得到结果。

- 入栈操作： 按照压栈序列的顺序执行。
- 出栈操作： 每次入栈后，循环判断 “栈顶元素 == 弹出序列的当前元素” 是否成立，将符合弹出序列顺序的栈顶元素全部弹出。

由于题目规定 栈的所有数字均不相等 ，因此在循环入栈中，每个元素出栈的位置的可能性是唯一的（若有重复数字，则具有多个可出栈的位置）。因而，在遇到 “栈顶元素 == 弹出序列的当前元素” 就应立即执行出栈。

### 代码

---

```go
func validateStackSequences(pushed []int, popped []int) bool {
    var stack []int
    length := len(pushed) + len(popped)
    if length == 0 {
        return true
    }
    for i := 0; i < length; i++ {
        if len(pushed) > 0 {
            stack = append(stack, pushed[0])
            pushed = pushed[1:]
        }
        for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
            stack = stack[:len(stack)-1]
            popped = popped[1:]
            if len(popped) == 0 {
                return true
            }
        }
    }

    return false
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)，n = len(pushed) + len(popped)。

- 空间复杂度

    O(n)，使用了一个辅助的栈来模拟。