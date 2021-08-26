package T28_isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/************* 递归法 **************/
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		return recur(root.Left, root.Right)
	}
}

func recur(L, R *TreeNode) bool {
	if L == nil && R == nil {
		return true
	} else if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return recur(L.Left, R.Right) && recur(L.Right, R.Left)
}

/************** 迭代法 ****************/
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var leftStack []*TreeNode
	leftStack = append(leftStack, root.Left)
	var rightStack []*TreeNode
	rightStack = append(rightStack, root.Right)

	var leftNode *TreeNode
	var rightNode *TreeNode

	for len(leftStack) != 0 || len(rightStack) != 0 {
		leftNode = leftStack[len(leftStack)-1]
		leftStack = leftStack[:len(leftStack)-1]
		rightNode = rightStack[len(rightStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		if leftNode == nil && rightNode == nil {
			continue
		} else if leftNode != nil && rightNode != nil {
			leftStack = append(leftStack, leftNode.Left)
			leftStack = append(leftStack, leftNode.Right)
			rightStack = append(rightStack, rightNode.Right)
			rightStack = append(rightStack, rightNode.Left)
		} else {
			return false
		}

		if rightNode.Val != leftNode.Val {
			return false
		}
	}

	return true
}
