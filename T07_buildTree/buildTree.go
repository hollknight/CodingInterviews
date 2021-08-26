package T07_buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]

	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}

	if index > 0 {
		root.Left = buildTree(preorder[1:index+1], inorder[:index])
	} else {
		root.Left = buildTree(nil, nil)
	}

	root.Right = buildTree(preorder[index+1:], inorder[index+1:])

	return root
}
