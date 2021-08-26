package T26_isSubStructure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return is(A, B, true)
}

func is(A *TreeNode, B *TreeNode, isRoot bool) bool {
	if A == nil || B == nil {
		if A == B {
			return true
		} else {
			return false
		}
	}

	if A.Val == B.Val {
		if B.Left == nil {
			return is(A.Right, B.Right, false) || is(A.Left, B, true) || is(A.Right, B, true)
		} else if B.Right == nil {
			return is(A.Left, B.Left, false) || is(A.Left, B, true) || is(A.Right, B, true)
		} else {
			return is(A.Left, B.Left, false) && is(A.Right, B.Right, false) || is(A.Left, B, true) || is(A.Right, B, true)
		}
	} else {
		if isRoot {
			return is(A.Left, B, true) || is(A.Right, B, true)
		} else {
			return false
		}
	}
}
