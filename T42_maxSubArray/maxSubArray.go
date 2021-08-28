package T42_maxSubArray

func maxSubArray(nums []int) int {
	max := nums[0]
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if temp > max {
			if temp == 0 {
				if nums[i] == 0 {
					max = temp
				}
				continue
			}
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
	}

	return max
}
