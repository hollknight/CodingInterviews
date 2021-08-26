# 用两个栈实现队列

- 标签: 栈, 队列
- 链接: https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
- 难度: 简单
- 题号: 09

### 题目描述

---

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

**示例1：**

```
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
```

**示例2：**

```
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
```

### 思路

---

使用两个栈，`in`和`out`，当`appendTail`时，直接`push`进`in`栈中。当`deleteHead`时，如果`out`栈中有元素，就按次序`pop`出来，如果`out`栈为空，则按次序将`in`栈中所有元素`pop`出来，`push`进`out`栈中，再从`out`栈中`pop`出元素，如果两个栈都为空，则`return -1`。

![https://assets.leetcode-cn.com/solution-static/jianzhi_09/jianzhi_9.gif](https://assets.leetcode-cn.com/solution-static/jianzhi_09/jianzhi_9.gif)

### 代码

---

```go
type CQueue struct {
    in stack
    out stack
}

type stack []int

func (s *stack) push(num int) {
    *s = append(*s, num)
}

func (s *stack) pop() int {
    if len(*s) == 0 {
        return -1
    } else {
        num := (*s)[len(*s)-1]
        *s = (*s)[:len(*s)-1]
        return num
    }
}

func Constructor() CQueue {
    var in stack
    var out stack
    queue := CQueue{
        in: in,
        out: out,
    }

    return queue
}

func (this *CQueue) AppendTail(value int)  {
    this.in.push(value)
}

func (this *CQueue) DeleteHead() int {
    if len(this.out) == 0 {
        if len(this.in) == 0 {
            return -1
        } else {
            len := len(this.in)
            for i := 0; i < len; i++ {
                this.out.push(this.in.pop())
            }
        }
    }
    return this.out.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
```

### 复杂度分析

---

- 时间复杂度

  每次操作均为O(1)。

- 空间复杂度

  O(n)，使用两个栈。