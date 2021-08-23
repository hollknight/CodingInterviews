# 表示数值的字符串

- 标签: 字符串
- 链接: https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
- 难度: 中等
- 题号: 20

### 题目描述

---

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

**数值（按顺序）可以分成以下几个部分：**

1. 若干空格
2. 一个 小数 或者 整数
3. （可选）一个 'e' 或 'E' ，后面跟着一个 整数
4. 若干空格

**小数（按顺序）可以分成以下几个部分：**

1. （可选）一个符号字符（'+' 或 '-'）
2. 下述格式之一：
    - 至少一位数字，后面跟着一个点 '.'
    - 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
    - 一个点 '.' ，后面跟着至少一位数字

**整数（按顺序）可以分成以下几个部分：**

1. （可选）一个符号字符（'+' 或 '-'）
2. 至少一位数字

**示例1：**

```
部分数值列举如下：
["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：
["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]
```

### 思路

---

运用`go`语言中将`string`转换为整数与小数的库函数，并对`error`做判断，如果`error`不为`nil`，则证明字符串不能转换为数字。

有限状态机的方法见题解：

[](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/mian-shi-ti-20-biao-shi-shu-zhi-de-zi-fu-chuan-y-2/)

### 代码

---

```go
func isNumber(s string) bool {
    s = strings.TrimSpace(s)
    var strList []string
    strList = append(strList, s)
    if find := strings.Contains(s, "e"); find {
        strList = strings.Split(s, "e")
    } else if find := strings.Contains(s, "e"); find {
        strList = strings.Split(s, "E")
    }
    if len(strList) > 2 {
        return false
    } else if len(strList) == 2 {
        if _, err := strconv.Atoi(strList[1]); err != nil {
            return false
        }
    }
    if _, err := strconv.ParseFloat(strList[0], 32); err != nil {
        return false
    }
    return true
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(n)。