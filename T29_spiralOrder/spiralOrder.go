package T29_spiralOrder

func spiralOrder(matrix [][]int) []int {
	var nums []int
	rowBegin := 0
	rowEnd := len(matrix) - 1
	if rowEnd == -1 {
		return nums
	}
	columBegin := 0
	columEnd := len(matrix[0]) - 1

	for columBegin <= columEnd && rowBegin <= rowEnd {
		for i := columBegin; i <= columEnd; i++ {
			nums = append(nums, matrix[rowBegin][i])
		}
		rowBegin++
		for i := rowBegin; i <= rowEnd; i++ {
			nums = append(nums, matrix[i][columEnd])
		}
		columEnd--
		for i := columEnd; i >= columBegin; i-- {
			nums = append(nums, matrix[rowEnd][i])
		}
		rowEnd--
		for i := rowEnd; i >= rowBegin; i-- {
			nums = append(nums, matrix[i][columBegin])
		}
		columBegin++
	}
	nums = nums[:len(matrix)*len(matrix[0])]

	return nums
}
