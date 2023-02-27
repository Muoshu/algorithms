package design

import (
	"algorithms/mysort"
	"fmt"
	"sort"
)

func MaxSubsequenceSumDP(x []int) {
	c, sum, b := 0, 0, 0
	for i := 0; i < len(x); i++ {
		if b > 0 {
			b += x[i]
		} else {
			b = x[i]
		}
		if b > sum {
			sum = b
			c = i
		}
	}
	fmt.Println(sum, c)
}
func MaxSubsequenceSumEnumerate(x []int) {
	sum, first, last := 0, 0, 0
	n := len(x)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			thisSum := 0
			for k := i; k <= j; k++ {
				thisSum += x[k]
				if thisSum > sum {
					sum = thisSum
					first, last = i, j
				}
			}
		}
	}
	fmt.Println(first, last)
}

func FindMinMax(x []int) (min int, max int) {
	n := len(x)
	if n == 0 {
		panic("Slice nums cannot be 0-size.")
	}
	min, max = x[0], x[0]
	for i := 0; i < n; i++ {
		if min > x[i] {
			min = x[i]
		}
		if max < x[i] {
			max = x[i]
		}
	}
	return min, max
}
func FindMinMax2(x []int) (min int, max int) {
	n := len(x)
	var t int
	if n%2 == 1 {
		t = n - 1
	} else {
		t = n
	}
	if n == 0 {
		panic("Slice nums cannot be 0-size.")
	}
	for i := 0; i < t; i += 2 {
		if x[i] > x[i+1] {
			x[i], x[i+1] = x[i+1], x[i]
		}
	}
	min, max = x[0], x[0]
	for i := 0; i < t; i += 2 {
		if min > x[i] {
			min = x[i]
		}
		if max < x[i+1] {
			max = x[i+1]
		}
	}
	if min > x[n-1] {
		min = x[n-1]
	}
	if max < x[n-1] {
		max = x[n-1]
	}
	return min, max
}
func FindSecond(x []int) int {
	n := len(x)
	if n == 0 {
		panic("Slice nums cannot be 0-size.")
	}
	max, secondMax := x[0], x[0]
	for i := 0; i < n; i++ {
		if secondMax < x[i] {
			if x[i] <= max {
				secondMax = x[i]
			} else {
				secondMax, max = max, x[i]
			}

		}
	}
	return secondMax
}

func Championship(x []int) int {
	n := len(x)
	if n == 0 {
		panic("Slice nums cannot be 0-size.")
	}
	loser := make([][]int, n)
	k := n
	var m int
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	for k >= 2 {

		j := 0
		if k%2 == 1 {
			m = k/2 + 1
		} else {
			m = k / 2
		}
		nextTurn := make([]int, m)
		for i := 0; i/2 < k/2; i += 2 {
			a, b := idx[i], idx[i+1]
			if x[a] > x[b] {
				loser[a] = append(loser[a], x[b])
				nextTurn[j] = a
				j++
			} else {
				loser[b] = append(loser[b], x[a])
				nextTurn[j] = b
				j++
			}
		}
		if k%2 == 1 {
			nextTurn[j] = idx[k-1]
		}
		k = m
		for i := 0; i < m; i++ {
			idx[i] = nextTurn[i]
		}
	}
	maxIndex := idx[0]

	return FindMax(loser[maxIndex])
}
func FindMax(x []int) int {
	n := len(x)
	if n == 0 {
		panic("Slice can not be 0 size.")
	}
	max := x[0]
	for i := 0; i < n; i++ {
		if max < x[i] {
			max = x[i]
		}
	}
	return max
}

func KthMin(x []int, k int) int {
	n := len(x)
	mStar := FindMid(x)
	var s1, s2 []int
	for i := 0; i < n; i++ {
		if x[i] < mStar {
			s1 = append(s1, x[i])
		}
		if x[i] > mStar {
			s2 = append(s2, x[i])
		}
	}
	if k == len(s1)+1 {
		return mStar
	}
	if k <= len(s1) {
		return KthMin(s1, k)
	} else {
		return KthMin(s2, k-len(s1)-1)
	}

}

// FindMid return mStar
func FindMid(x []int) int {
	n := len(x)
	var median []int
	if n == 0 {
		panic("Slice size cannot be 0.")
	}
	if n < 5 {
		sort.Ints(x)
		return x[n/2]
	}
	for i := 0; i+4 < n; i += 5 {
		sort.Ints(x[i : i+5])
		median = append(median, x[i : i+5][2])
	}
	sort.Ints(median)
	s := len(median)
	return median[s/2]
}

func SelectK(x []int, low, high, k int) int {
	pivot := mysort.Partition(x, low, high)
	rank := pivot - low + 1 //枢轴到左边的距离
	fmt.Println(pivot, rank)
	fmt.Println(x)
	if rank == k {
		return x[pivot]
	}
	if rank > k {
		return SelectK(x, low, pivot-1, k)
	}
	return SelectK(x, pivot+1, high, k-rank)
}
