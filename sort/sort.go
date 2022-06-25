package sort

// InsertSort implement intType insert sort
func InsertSort(x []int) []int {
	for i := 1; i < len(x); i++ {
		temp := x[i]
		for j := i - 1; j >= 0; j-- {
			if temp < x[j] {
				x[j+1], x[j] = x[j], temp
			}
		}

	}
	return x
}

func SplitInsertSort(x []int) []int {
	for i := 1; i < len(x); i++ {
		temp := x[i]
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) / 2
			if x[mid] > temp {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j := i - 1; j >= high+1; j-- {
			x[j+1] = x[j]
		}
		x[high+1] = temp
	}
	return x
}

func ShellSort(x []int) []int {
	for d := len(x) / 2; d > 0; d = d / 2 {
		for i := d + 1; i < len(x); i++ {
			if x[i] < x[i-d] {
				temp := x[i]
				j := i - d
				for ; j > 0 && temp < x[j]; j = j - d {
					x[j+d] = x[j]
				}
				x[j+d] = temp
			}
		}
	}
	return x
}

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

func QuickSort(x []int, low, high int) {
	if low < high {
		pivot := partition(x, low, high)
		QuickSort(x, low, pivot-1)
		QuickSort(x, pivot+1, high)
	}

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

func partition(x []int, low, high int) int {
	pivot := x[low]
	for low < high {
		for pivot <= x[high] && low < high {
			high--
		}
		x[low] = x[high]
		for pivot > x[low] && low < high {
			low++
		}
		x[high] = x[low]
	}
	x[low] = pivot //枢轴最终位置
	return low
}

func HeapSort(x []int) {

	BuildMaxHeap(x)
	for i := len(x) - 1; i > 0; i-- {
		x[i], x[0] = x[0], x[i]
		heapAdjust(x, 0, i)
	}

}

func BuildMaxHeap(x []int) {
	for i := len(x)/2 - 1; i >= 0; i-- {
		heapAdjust(x, i, len(x))
	}
}

func heapAdjust(x []int, k, heapSize int) {
	l := 2*k + 1 //左子树下标
	r := l + 1   //右子树下标

	largest := k
	if l < heapSize && x[l] > x[largest] {
		largest = l
	}
	if r < heapSize && x[r] > x[largest] {
		largest = r
	}
	if k != largest {
		x[k], x[largest] = x[largest], x[k]
		heapAdjust(x, largest, heapSize)
	}
}
