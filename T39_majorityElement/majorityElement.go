package T39_majorityElement

func majorityElement(nums []int) int {
	vote := 0
	majority := 0
	for i := 0; i < len(nums); i++ {
		if vote == 0 {
			majority = nums[i]
		}
		if majority == nums[i] {
			vote++
		} else {
			vote--
		}
	}

	return majority
}
