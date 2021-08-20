package _4_kthLargest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || k <= 0 {
			return
		}
		dfs(root.Right)
		k--
		if k == 0 {
			ans = root.Val
			return
		}
		dfs(root.Left)
	}
	dfs(root)
	return ans
}
