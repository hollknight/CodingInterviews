package _1_exchange

func exchange(nums []int) []int {
	begin := 0
	end := len(nums) - 1
	for begin < end {
		for nums[begin]%2 == 1 && begin < end {
			begin++
		}
		for nums[end]%2 == 0 && begin < end {
			end--
		}
		nums[begin], nums[end] = nums[end], nums[begin]
	}

	return nums
}
