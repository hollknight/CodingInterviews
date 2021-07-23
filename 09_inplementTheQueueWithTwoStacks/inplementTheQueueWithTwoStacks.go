package _9_inplementTheQueueWithTwoStacks

type CQueue struct {
	in  stack
	out stack
}

type stack []int

func (s *stack) push(num int) {
	*s = append(*s, num)
}

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	} else {
		num := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return num
	}
}

func Constructor() CQueue {
	var in stack
	var out stack
	queue := CQueue{
		in:  in,
		out: out,
	}

	return queue
}

func (this *CQueue) AppendTail(value int) {
	this.in.push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		if len(this.in) == 0 {
			return -1
		} else {
			len := len(this.in)
			for i := 0; i < len; i++ {
				this.out.push(this.in.pop())
			}
		}
	}
	return this.out.pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
