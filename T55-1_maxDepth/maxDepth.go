package T55_1_maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			maxDepth = depth
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 1)
	return maxDepth
}
