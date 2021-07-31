package _1_minArray

func minArray(numbers []int) int {
	begin := 0
	end := len(numbers) - 1
	mid := (begin + end) / 2

	for begin <= end {
		// 如果最左端元素小于最右端元素，则此时左到右递增，最小元素即为最左端元素
		if numbers[begin] < numbers[end] {
			return numbers[begin]
		}
		// 如果中心元素大于最左端元素，则中心元素一定在最小值左侧，将最左端移到mid+1处
		if numbers[mid] > numbers[begin] {
			begin = mid + 1
			// 如果中心元素小于最左端元素，则中心元素在最小值处或最小值右侧，将最右端移到mid处
		} else if numbers[mid] < numbers[begin] {
			end = mid
			// 此时中心元素等于最左侧元素，将最左端右移
		} else {
			begin++
		}
		mid = (begin + end) / 2
	}
	mid = (begin + end) / 2

	return numbers[mid]
}
