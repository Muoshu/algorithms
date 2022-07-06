package design

import "fmt"

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
			m = 2/k + 1
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
