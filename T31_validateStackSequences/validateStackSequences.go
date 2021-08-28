package T31_validateStackSequences

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	length := len(pushed) + len(popped)
	if length == 0 {
		return true
	}
	for i := 0; i < length; i++ {
		if len(pushed) > 0 {
			stack = append(stack, pushed[0])
			pushed = pushed[1:]
		}
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
			if len(popped) == 0 {
				return true
			}
		}
	}

	return false
}
