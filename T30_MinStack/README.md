# 包含min函数的栈

- 标签: 栈
- 链接: https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/
- 难度: 简单
- 题号: 30

### 题目描述

---

定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

**示例1：**

```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.
```

### 思路

---

定义一个辅助栈，辅助栈只储存非严格递减的数值。

### 代码

---

```go
type MinStack struct {
    stack []int
    minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minStack: []int{},
    }
}

func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, x)
    }
}

func (this *MinStack) Pop()  {
    if this.Top() == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
    return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
```

### 复杂度分析

---

- 时间复杂度

    均为O(1)。

- 空间复杂度

    O(n)，辅助栈。