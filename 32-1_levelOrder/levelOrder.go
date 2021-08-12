package _2_1_levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var ans []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			ans = append(ans, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return ans
}
