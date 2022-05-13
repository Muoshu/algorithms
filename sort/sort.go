package sort

func BubbleSort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}

func SelectSort(x []int) {
	for i := 0; i < len(x); i++ {
		temp := x[i]
		for j := 0; j < len(x); j++ {
			if temp < x[j] {
				x[i], x[j] = x[j], x[i]
				temp = x[j]
			}
		}
	}
}

func MergeSort(x []int) []int {
	if len(x) <= 1 {
		return x
	}
	mid := len(x) / 2
	left := MergeSort(x[:mid])
	right := MergeSort(x[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) >= 0 || len(right) >= 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
