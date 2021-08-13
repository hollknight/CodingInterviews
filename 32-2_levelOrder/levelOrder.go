package _2_2_levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	queue = append(queue, root)
	var ans [][]int

	for len(queue) != 0 {
		var row []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
				row = append(row, node.Val)
			}
		}
		ans = append(ans, row)
	}

	return ans[:len(ans)-1]
}
