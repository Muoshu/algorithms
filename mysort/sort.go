package mysort

import "sort"

// InsertSort implement intType insert sort
func InsertSort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := i; j > 0 && x[j-1] > x[j]; j-- {
			x[j-1], x[j] = x[j], x[j-1]
		}
	}
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
	var swapped bool
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

func SelectSort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		min := i
		for j := i + 1; j < len(x); j++ {
			if x[min] > x[j] {
				min = j
			}
		}
		if min != i {
			x[i], x[min] = x[min], x[i]
		}
	}
}

func MergeSort(x []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		MergeSort(x, low, mid)
		MergeSort(x, mid+1, high)
		merge(x, low, mid, high)
	}

}

func merge(x []int, low, mid, high int) {
	B := make([]int, len(x))
	for k := low; k <= high; k++ {
		B[k] = x[k]
	}
	i, j, k := low, mid+1, low
	for ; i <= mid && j <= high; k++ {
		if B[i] <= B[j] {
			x[k] = B[i]
			i++
		} else {
			x[k] = B[j]
			j++
		}
	}
	for i <= mid { // 第一个表未检测完，复制；两个for只会有一个执行
		x[k] = B[i]
		i++
		k++
	}
	for j <= high { // 第二表未检测完，复制
		x[k] = B[j]
		k++
		j++
	}
}

func QuickSort(x []int, low, high int) {
	if low < high {
		pivot := Partition(x, low, high)
		QuickSort(x, low, pivot-1)
		QuickSort(x, pivot+1, high)
	}

}

func Partition(x []int, low, high int) int {
	pivot := x[low]
	for low < high {
		for low < high && x[high] >= pivot { //找到右边第一个比pivot小
			high--
		}
		x[low] = x[high]
		for low < high && pivot >= x[low] { //找到左边第一个比pivot大
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
	size := len(x)
	for i := len(x)/2 - 1; i >= 0; i-- {
		heapAdjust(x, i, size)
	}
}

func heapAdjust(x []int, k, heapSize int) {
	l := k<<1 + 1 //左子树下标
	r := k<<1 + 2 //右子树下标
	largest := k
	if l < heapSize && x[l] > x[largest] {
		largest = l
	}
	if r < heapSize && x[r] > x[largest] {
		largest = r
	}
	if largest != k {
		x[largest], x[k] = x[k], x[largest]
		heapAdjust(x, largest, heapSize)
	}
}

func BucketSort(x []int) {
	n := len(x)
	if n <= 0 {
		return
	}
	mn, mx := x[0], x[0]
	for _, val := range x {
		mn = Min(val, mn)
		mx = Max(val, mx)
	}
	size := (mx-mn)/n + 1   //桶大小
	cnt := (mx-mn)/size + 1 //桶个数
	buckets := make([][]int, cnt)
	for _, val := range x {
		idx := (val - mn) / size
		buckets[idx] = append(buckets[idx], val)
	}
	//对各桶中的数据排序
	for i := 0; i < cnt; i++ {
		sort.Ints(buckets[i])
	}
	var index int
	for i := 0; i < cnt; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			x[index] = buckets[i][j]
			index++
		}
	}
}

func Min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
