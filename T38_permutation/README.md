# 字符串的排列

- 标签: dfs
- 链接: https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
- 难度: 中等
- 题号: 38

### 题目描述

---

输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

**示例1：**

```
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]
```

### 思路

---

[力扣](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/)

**终止条件：** 当 x = len(c) - 1 时，代表所有位已固定（最后一位只有 1 种情况），则将当前组合 c 转化为字符串并加入 res ，并返回；
**递推参数：** 当前固定位 x ；
**递推工作：** 初始化一个 Set ，用于排除重复的字符；将第 x 位字符与 i ∈ [x, len(c)] 字符分别交换，并进入下层递归；
**剪枝：** 若 c[i] 在 Set 中，代表其是重复字符，因此 “剪枝” ；
将 c[i] 加入 Set ，以便之后遇到重复字符时剪枝；
**固定字符：** 将字符 c[i] 和 c[x] 交换，即固定 c[i] 为当前位字符；
**开启下层递归：** 调用 dfs(x + 1) ，即开始固定第 x + 1 个字符；
**还原交换：** 将字符 c[i] 和 c[x] 交换（还原之前的交换）；

### 代码

---

```go
func permutation(s string) []string {
	res := []string{}
	bytes := []byte(s)
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(bytes)-1 {
			res = append(res, string(bytes))
			return
		}
		dict := map[byte]bool{}
		for i := x; i < len(bytes); i++ {
			if !dict[bytes[i]] {
				bytes[x], bytes[i] = bytes[i], bytes[x]
				dict[bytes[x]] = true
				dfs(x + 1)
				bytes[x], bytes[i] = bytes[i], bytes[x]
			}
		}
	}
	dfs(0)
	return res
}
```

### 复杂度分析

---

- 时间复杂度

    O(n*(n!))。

- 空间复杂度

    O(n^2)。