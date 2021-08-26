package T27_mirrorTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorRoot := new(TreeNode)
	mirrorRoot.Val = root.Val
	mirrorRoot.Left = mirrorTree(root.Right)
	mirrorRoot.Right = mirrorTree(root.Left)

	return mirrorRoot
}
