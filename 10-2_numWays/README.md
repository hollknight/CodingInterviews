# 青蛙跳台阶问题

上次复习时间: 2021年7月29日 23:38
备注: 斐波那契数列
来源: 剑指 Offer
标签: 动态规划
记录时间: 2021年7月29日 23:11
链接: https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
难度: 简单
题号: 10- II

### 题目描述

---

一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

**示例1：**

```text
输入：n = 2
输出：2
```

**示例2：**

```text
输入：n = 0
输出：1
```

### 思路

---

![https://pic.leetcode-cn.com/108249e4d62d429f9cd6cab5bbd6afca581ee61c7d762a4c8ea0c62e08e10762-Picture13.png](https://pic.leetcode-cn.com/108249e4d62d429f9cd6cab5bbd6afca581ee61c7d762a4c8ea0c62e08e10762-Picture13.png)

做法和斐波那契数列相同。

### 代码

---

```go
func numWays(n int) int {
    ppre := 1
    pre := 1
    ans := 1
    for i := 2; i <= n; i++ {
        ans = ppre + pre
        if ans >= 1e9+7 {
            ans %= 1e9+7
        }
        ppre = pre
        pre = ans
    }

    return ans
}
```

### 复杂度分析

---

- 时间复杂度

    O(n)。

- 空间复杂度

    O(1)。