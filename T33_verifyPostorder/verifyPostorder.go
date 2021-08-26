package T33_verifyPostorder

import "math"

func verifyPostorder(postorder []int) bool {
	var stack []int
	root := math.MaxInt32
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}

	return true
}
