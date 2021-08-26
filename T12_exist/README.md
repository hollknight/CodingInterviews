# 矩阵中的路径

- 标签: dfs
- 链接: https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
- 难度: 中等
- 题号: 12

### 题目描述

---

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

![https://assets.leetcode.com/uploads/2020/11/04/word2.jpg](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)

**示例1：**

```text
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
```

### 思路

---

本问题是典型的矩阵搜索问题，可使用 **深度优先搜索（DFS）+ 剪枝** 解决。

- **深度优先搜索**： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
- **剪枝：** 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。

    ![https://pic.leetcode-cn.com/1604944042-glmqJO-Picture0.png](https://pic.leetcode-cn.com/1604944042-glmqJO-Picture0.png)

**DFS 解析：**

- **递归参数：** 当前元素在矩阵 board 中的行列索引 i 和 j ，当前目标字符在 word 中的索引 k 。
- **终止条件：**
    1. 返回 falsefalse ： (1) 行或列索引越界 或 (2) 当前矩阵元素与目标字符不同 或 (3) 当前矩阵元素已访问过 （ (3) 可合并至 (2) ） 。
    2. 返回 truetrue ： k = len(word) - 1 ，即字符串 word 已全部匹配。
- **递推工作：**
    1. 标记当前矩阵元素： 将 board[i][j] 修改为 空字符 '\' ，代表此元素已访问过，防止之后搜索时重复访问。
    2. 搜索下一单元格： 朝当前元素的 上、下、左、右 四个方向开启下层递归，使用 或 连接 （代表只需找到一条可行路径就直接返回，不再做后续 DFS ），并记录结果至 res 。
    3. 还原当前矩阵元素： 将 board[i][j] 元素还原至初始值，即 word[k] 。
- **返回值：** 返回布尔量 res ，代表是否搜索到目标字符串。

### 代码

---

```go
func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if findWay(board, word, i, j) {
                return true
            }
        }
    }
    return false
}

func findWay(board [][]byte, word string, i, j int) bool {
    if len(word) == 0 {
        return true
    }
    if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 || board[i][j] != word[0] {
        return false
    }

		// 将已判断元素zhikong
    board[i][j] = '0'

    res := findWay(board, word[1:], i+1, j) || findWay(board, word[1:], i-1, j) ||
           findWay(board, word[1:], i, j+1) || findWay(board, word[1:], i, j-1)
    board[i][j] = word[0]

    return res
}
```

### 复杂度分析

---

假设`borad`的行数和列数分别为`m`、`n`，`word`长度为k。

- 时间复杂度

    O(3^k*mn)，最差情况下，需要遍历矩阵中长度为 k 字符串的所有方案，时间复杂度为 O(3^k)；矩阵中共有 mn 个起点，时间复杂度为 O(mn) 。

    - 方案数计算：

    设字符串长度为 k ，搜索中每个字符有上、下、左、右四个方向可以选择，舍弃回头（上个字符）的方向，剩下 33 种选择，因此方案数的复杂度为 O(3^k)。

- 空间复杂度

    O(k)，搜索过程中的递归深度不超过 k ，因此系统因函数调用累计使用的栈空间占用 O(k) （因为函数返回后，系统调用的栈空间会释放）。最坏情况下 k = mn ，递归深度为 mn ，此时系统栈使用 O(mn) 的额外空间。