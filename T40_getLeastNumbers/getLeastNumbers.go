package T40_getLeastNumbers

func getLeastNumbers(arr []int, k int) []int {
	quickSort(&arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSort(arr *[]int, left, right, k int) {
	if left >= right {
		return
	}
	i := left
	j := right
	for i < j {
		for i < j && (*arr)[j] >= (*arr)[left] {
			j--
		}
		for i < j && (*arr)[i] <= (*arr)[left] {
			i++
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
	(*arr)[left], (*arr)[i] = (*arr)[i], (*arr)[left]
	if i > k {
		quickSort(arr, left, i-1, k)
	} else if i < k {
		quickSort(arr, i+1, right, k)
	}
}
