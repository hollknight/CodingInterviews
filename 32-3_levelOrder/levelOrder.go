package _2_3_levelOrder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	queue = append(queue, root)
	var ans [][]int
	rowFlag := 0

	for len(queue) != 0 {
		var row []int
		for i := len(queue); i > 0; i-- {
			var node *TreeNode
			node = queue[0]
			queue = queue[1:]

			if node != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
				row = append(row, node.Val)
			}
		}
		if rowFlag%2 == 1 {
			row = reverse(row)
		}
		rowFlag++
		ans = append(ans, row)
	}

	return ans[:len(ans)-1]
}

func reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
