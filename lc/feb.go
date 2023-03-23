package lc

import (
	"algorithms/Link"
	"algorithms/binaryTree"
	stack2 "algorithms/stack"
	"fmt"
	"sort"
	"strings"
)

func ReplaceWords(dictionary []string, sentence string) string {
	sentenceWords := strings.Split(sentence, " ")
	m := len(dictionary)
	n := len(sentenceWords)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if strings.HasPrefix(dictionary[j], dictionary[i]) {
				dictionary[j] = dictionary[i]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if strings.HasPrefix(sentenceWords[j], dictionary[i]) {
				sentenceWords[j] = dictionary[i]
			}
		}
	}
	return strings.Join(sentenceWords, " ")
}

func FindNumNotApp(arr []int) []int {
	var s = len(arr)
	var app = make([]int, s+1)
	for i := 0; i < s; i++ {
		temp := arr[i]
		app[temp] = temp
	}
	var ans []int
	for key, val := range app {
		if key == 0 {
			continue
		}
		if val == 0 {
			ans = append(ans, key)
		}
	}
	return ans
}

func SearchMatrix(a [][]int, target int) bool {
	mLength := len(a[0])
	i, j := 0, mLength-1
	for i < mLength && j >= 0 {
		if a[i][j] == target {
			return true
		} else if a[i][j] < target {
			i++
		} else {
			j--
		}

	}
	return false
}

func RotateMatrix(a [][]int) [][]int {
	mLength := len(a[0])
	for i := 0; i < mLength; i++ {
		for j := i; j < mLength; j++ {
			temp := a[i][j]
			a[i][j] = a[j][i]
			a[j][i] = temp
		}
	}
	for j := 0; j < mLength/2; j++ {
		for i := 0; i < mLength; i++ {
			temp := a[i][j]
			a[i][j] = a[i][mLength-j-1]
			a[i][mLength-j-1] = temp

		}
	}
	return a
}

func ValidPar(s string) bool {
	sLength := len(s)
	stack := stack2.NewStack()
	if sLength%2 != 0 {
		return false
	}
	for i := 0; i < sLength; i++ {
		if string(s[i]) == "{" || string(s[i]) == "[" || string(s[i]) == "(" {
			stack.Push(string(s[i]))
		} else {
			if stack.IsEmpty() {
				return false
			}
			temp := stack.Pop()
			if string(s[i]) == "}" && temp != "{" {
				return false
			}
			if string(s[i]) == "]" && temp != "[" {
				return false
			}
			if string(s[i]) == ")" && temp != "(" {
				return false
			}
		}
	}
	return true
}

func DailyTem(a []int) []int {
	aLength := len(a)
	ans := make([]int, aLength)
	stack := stack2.NewStack()

	for i := 0; i < aLength; i++ {
		for !stack.IsEmpty() {
			if a[i] > a[stack.Peek().(int)] {
				ans[stack.Peek().(int)] = i - stack.Peek().(int)
				stack.Pop()
			} else {
				break
			}
		}
		stack.Push(i)
	}

	//for i := 0; i < aLength-1; i++ {
	//	if a[i] > a[i+1] {
	//		if i > 0 && stack.Peek() != nil {
	//			if a[i] > a[stack.Peek().(int)] {
	//				ans[stack.Peek().(int)] = i - stack.Peek().(int)
	//				fmt.Println()
	//				stack.Pop()
	//			}
	//
	//		}
	//		stack.Push(i)
	//	}
	//	if a[i] < a[i+1] {
	//		ans[i] = 1
	//	}
	//	if i > 0 && stack.Peek() != nil {
	//		if a[i] > a[stack.Peek().(int)] {
	//			ans[stack.Peek().(int)] = i - stack.Peek().(int)
	//			stack.Pop()
	//		}
	//
	//	}
	//}
	return ans
}

func MergeKSortedList(a []Link.SingleLink) *Link.SingleLink {
	var maxSize int
	var ans = Link.NewSingleLink()
	var heap = binaryTree.NewHeap[int]()

	for maxSize < len(a) {
		for i := 0; i < len(a); i++ {
			if x := a[i].RemoveAtBeg(); x != nil {
				heap.Push(x.(int))
			} else {
				maxSize++
			}
		}
		temp := len(a) - maxSize
		for temp > 0 {
			ans.Append(heap.Pop())
			temp--
		}
		if maxSize != len(a) {
			maxSize = 0
		}
	}
	return ans
}

func TowSum(a []int, target int) []int {
	m := make(map[int]int)
	for i, x := range a {
		if idx, ok := m[target-x]; ok {
			return []int{idx, i}
		}
		m[x] = i
	}
	fmt.Println(m)
	return nil
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, next := head, head.Next
	for next != nil {
		temp := curr.Val
		curr.Val = next.Val
		next.Val = temp

		curr = next.Next
		next = next.Next.Next
	}
	return head
}

func EraseOverlapIntervals(intervals [][]int) int {
	//intervals := [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}
	less := func(i, j int) bool {
		//if intervals[i][0] == intervals[j][0] {
		//	return intervals[i][0] < intervals[j][0]
		//}
		return intervals[i][1] < intervals[j][1]
	}

	sort.Slice(intervals, less)

	fmt.Println(intervals)
	temp := intervals[0][1]
	var count int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < temp {
			count++
		} else {
			temp = intervals[i][1]
			fmt.Println(i)
		}
	}
	return count
}

func TwoSumDoublePointer(nums []int, target int) []int {
	size := len(nums)
	if size <= 1 {
		return nil
	}
	head, tail := 0, size-1
	for head != tail {
		if nums[head]+nums[tail] < target {
			head++
		} else if nums[head]+nums[tail] > target {
			tail--
		} else {
			return []int{head + 1, tail + 1}
		}
	}
	return nil
}

func MinWindowSubstring(s, t string) string {
	sLen := len(s)
	tLen := len(t)
	if s == "" || t == "" || sLen < tLen {
		return ""
	}
	var winFreq [128]int
	var tFreq [128]int

	for i := 0; i < tLen; i++ {
		tFreq[t[i]]++
	}
	distance, minLen, begin, left, right := 0, sLen+1, 0, 0, 0
	for right < sLen {
		if tFreq[s[right]] == 0 {
			right++
			continue
		}
		if winFreq[s[right]] < tFreq[s[right]] {
			distance++
		}
		winFreq[s[right]]++
		right++
		for distance == tLen {
			if (right - left) < minLen {
				minLen = right - left
				begin = left
			}
			if tFreq[s[left]] == 0 {
				left++
				continue
			}
			if winFreq[s[left]] == tFreq[s[left]] {
				distance--
			}
			winFreq[s[left]]--
			left++
		}
	}
	if minLen == sLen+1 {
		return ""
	}
	return s[begin : begin+minLen]
}

func KthLargest(x []int, k int) int {
	l, r, target := 0, len(x)-1, len(x)-k
	for l < r {
		mid := quickSelection(x, l, r)
		if mid == target {
			return x[target]
		} else if mid < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return x[l]
}

func quickSelection(x []int, l, r int) int {
	i, j := l+1, r
	for true {
		for i < r && x[i] <= x[l] {
			i++
		}
		for l < j && x[j] >= x[l] {
			j--
		}
		if i >= j {
			break
		}
		x[i], x[j] = x[j], x[i]
	}
	x[l], x[j] = x[j], x[l]

	return j
}

func KthLargestV2(x []int, k int) int {
	h := make([]int, k)
	copy(h, x[:k])
	for i := k/2 - 1; i >= 0; i-- {
		buildHeap(h, i)
	}
	for i := k; i < len(x); i++ {
		if x[i] <= h[0] {
			continue
		}
		h[0] = x[i]
		for i := k/2 - 1; i >= 0; i-- {
			buildHeap(h, i)
		}
	}
	return h[0]
}

func buildHeap(x []int, parent int) {
	l := parent<<1 + 1
	r := parent<<1 + 2
	smallest := parent
	if l < len(x) && x[smallest] > x[l] {
		smallest = l
	}
	if r < len(x) && x[smallest] > x[r] {
		smallest = r
	}
	if parent != smallest {
		x[parent], x[smallest] = x[smallest], x[parent]
		buildHeap(x, smallest)
	}

}

func TopKFrequent(x []int, k int) []int {
	mn, mx := x[0], x[0]
	for _, val := range x {
		mn = min(val, mn)
		mx = max(mx, val)
	}
	m := make(map[int]int)
	for _, val := range x {
		m[val]++
	}
	s := make([][]int, len(m))
	i := 0
	for key, val := range m {
		s[i] = append(s[i], val, key)
		i++
	}
	lessFun := func(i, j int) bool {
		return s[i][0] > s[j][0]
	}
	sort.Slice(s, lessFun)
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, s[i][1])
	}
	return ans
}

func FindFirstLastPos(nums []int, target int) []int {
	size := len(nums)
	if size <= 0 {
		return []int{-1, -1}
	}
	l := findPosLeft(nums, target)
	r := findPosRight(nums, target) - 1
	if l == size || nums[l] != target {
		return []int{-1, -1}
	}
	return []int{l, r}
}

func findPosLeft(x []int, target int) int {
	l, r := 0, len(x)
	pos := (l + r) / 2
	for l < r {
		if x[pos] >= target { //继续向左边走
			r = pos
		} else { //向右边走
			l = pos + 1
		}
		pos = (r + l) / 2
	}
	return l
}

func findPosRight(x []int, target int) int {
	l, r := 0, len(x)
	pos := (l + r) / 2
	for l < r {
		if x[pos] > target {
			r = pos
		} else {
			l = pos + 1
		}
		pos = (r + l) / 2
	}
	return l
}

func SearchRotate(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] { //无法判断哪个区间是增序
			l++
		} else if nums[mid] <= nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}

func SingleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			r = mid
		}
	}
	return nums[l]
}

func SQ(x int) int {
	if x <= 0 {
		return 0
	}
	l, r := 1, x

	for l <= r {
		mid := (l + r) / 2
		sqrt := x / mid
		if sqrt == mid {
			return sqrt
		} else if sqrt < mid {
			r = mid + 1
		} else {
			l = mid - 1
		}
	}
	return r
}

func RecurMaxIsland(grid [][]int) int {
	directions := []int{-1, 0, 1, 0, -1}
	m := len(grid)
	n := len(grid[0])
	var stack [][]int
	var area, localArea int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				stack = append(stack, []int{i, j})
				grid[i][j] = 0
				localArea = 1
				for len(stack) > 0 {
					pos := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					for k := 0; k < 4; k++ {
						x, y := pos[0]+directions[k], pos[1]+directions[k+1]
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
							localArea++
							grid[x][y] = 0
							stack = append(stack, []int{x, y})
						}
					}
				}
				area = max(area, localArea)
			}
		}
	}
	return area
}

func MaxAreaIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var maxArea int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs1(grid, i, j))
			}
		}
	}
	return maxArea
}

func dfs1(grid [][]int, r, c int) int {
	directions := []int{-1, 0, 1, 0, -1}
	if grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	var x, y int
	area := 1
	for i := 0; i < 4; i++ {
		x = r + directions[i]
		y = c + directions[i+1]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			area += dfs1(grid, x, y)
		}
	}
	return area
}

func dfs2(grid [][]int, r, c int) int {
	if r < 0 || r > len(grid) || c < 0 || c > len(grid[0]) || grid[r][c] == 0 {
		return 0
	}
	return 1 + dfs2(grid, r-1, c) + dfs2(grid, r+1, c) + dfs2(grid, r, c+1) + dfs2(grid, r, c-1)
}

func FriendCircle(grid [][]int) int {
	visited := make([]bool, len(grid))
	circle := 0
	n := len(grid)
	var stack []int
	for i := 0; i < n; i++ {
		if !visited[i] {
			stack = append(stack, i)
			visited[i] = true
		} else {
			continue
		}
		for len(stack) != 0 {
			s := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for i := 0; i < n; i++ {
				if s != i && grid[s][i] == 1 {
					if !visited[i] {
						stack = append(stack, i)
						visited[i] = true
					}
				}
			}
		}
		circle++
	}
	return circle
}

func PacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var ans [][]int
	m, n := len(matrix), len(matrix[0])
	canReachP := make([][]bool, m)
	canReachA := make([][]bool, m)
	for i := 0; i < m; i++ {
		canReachP[i] = make([]bool, n)
		canReachA[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfsPA(matrix, canReachP, i, 0)
		dfsPA(matrix, canReachA, i, n-1)
	}
	//for i := 0; i < m; i++ {
	//	fmt.Println(canReachP[i])
	//}
	//fmt.Println("****************")
	//for i := 0; i < m; i++ {
	//	fmt.Println(canReachA[i])
	//}

	for i := 0; i < n; i++ {
		dfsPA(matrix, canReachP, 0, i)
		dfsPA(matrix, canReachA, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachA[i][j] && canReachP[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

func dfsPA(matrix [][]int, canReach [][]bool, r, c int) {
	directions := []int{-1, 0, 1, 0, -1}
	if canReach[r][c] {
		return
	}
	canReach[r][c] = true
	var x, y int
	for i := 0; i < 4; i++ {
		x = r + directions[i]
		y = c + directions[i+1]
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[r][c] <= matrix[x][y] {
			dfsPA(matrix, canReach, x, y)
		}
	}
}

// Permutations 全排列
func Permutations(nums []int) [][]int {
	var ans [][]int
	//backtracking(nums, 0, &ans)
	var dfs func(num []int, level int)

	dfs = func(num []int, level int) {
		if level == len(nums)-1 {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ans = append(ans, temp)
			return
		}
		for i := level; i < len(nums); i++ {
			nums[i], nums[level] = nums[level], nums[i]
			dfs(nums, level+1)
			nums[i], nums[level] = nums[level], nums[i]

		}
	}
	dfs(nums, 0)
	return ans
}

func backtracking(nums []int, level int, ans *[][]int) {
	if level == len(nums)-1 {
		//temp := make([]int, len(nums))
		//copy(temp, nums)
		//fmt.Printf("%p\n", ans)
		*ans = append(*ans, nums)
		//fmt.Printf("%p\n", ans)
		//fmt.Printf("%p\n", nums)
		return
	}
	for i := level; i < len(nums); i++ {
		nums[i], nums[level] = nums[level], nums[i]
		backtracking(nums, level+1, ans)
		nums[i], nums[level] = nums[level], nums[i]
	}
}

func Combine(n, k int) [][]int {
	var ans [][]int
	comb := make([]int, k)
	var count int
	CombBacktracking(&ans, comb, &count, 1, n, k)
	return ans
}

func CombBacktracking(ans *[][]int, comb []int, count *int, pos, n, k int) {
	if *count == k {
		temp := make([]int, k)
		copy(temp, comb)
		*ans = append(*ans, temp)

		return
	}
	for i := pos; i <= n; i++ {
		comb[*count] = i
		*count = *count + 1
		CombBacktracking(ans, comb, count, i+1, n, k)
		*count--
	}

}

func ShortestBridge(grid [][]int) int {
	directions := []int{-1, 0, 1, 0, -1}
	m, n := len(grid), len(grid[0])
	var points [][]int
	var flipped bool
	for i := 0; i < m; i++ {
		if flipped {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfsIsland(&points, grid, m, n, i, j)
				flipped = true
				break
			}
		}
	}
	//fmt.Println(grid)
	var x, y, level int
	for len(points) != 0 {
		level++
		nPoints := len(points)
		for nPoints > 0 {
			pos := points[0]
			points = points[1:]
			for i := 0; i < 4; i++ {
				x, y = pos[0]+directions[i], pos[1]+directions[i+1]
				if x < 0 || x > m || y < 0 || y > n || grid[x][y] == 2 {
					continue
				}
				if grid[x][y] == 1 {
					return level
				}
				grid[x][y] = 2
				points = append(points, []int{x, y})

			}
			nPoints--
		}
	}
	return 0
}

func dfsIsland(points *[][]int, grid [][]int, m, n, i, j int) {
	if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 2 {
		return
	}
	if grid[i][j] == 0 {
		*points = append(*points, []int{i, j})
		return
	}
	grid[i][j] = 2
	dfsIsland(points, grid, m, n, i-1, j)
	dfsIsland(points, grid, m, n, i+1, j)
	dfsIsland(points, grid, m, n, i, j-1)
	dfsIsland(points, grid, m, n, i, j+1)

}

func WordSearch(word string, board [][]byte) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	flag := false
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				wordSearchBT(word, board, i, j, 0, &flag, &visited)
			}
		}
	}
	return flag
}

func wordSearchBT(word string, board [][]byte, i, j, pos int, flag *bool, visited *[][]bool) {

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if (*visited)[i][j] || *flag || board[i][j] != word[pos] {
		return
	}
	if pos == len(word)-1 {
		*flag = true
		return
	}

	(*visited)[i][j] = true
	wordSearchBT(word, board, i-1, j, pos+1, flag, visited)
	wordSearchBT(word, board, i+1, j, pos+1, flag, visited)
	wordSearchBT(word, board, i, j-1, pos+1, flag, visited)
	wordSearchBT(word, board, i, j+1, pos+1, flag, visited)
	(*visited)[i][j] = false
}

func WordLadder(beginWord string, endWord string, wordList []string) [][]string {
	wLen := len(wordList)
	visited := make([]bool, wLen)
	var ans [][]string
	var path []string
	path = append(path, beginWord)
	for len(path) > 0 {
		word := path[0]
		path = path[1:]
		for i, w := range wordList {
			if isDiffer1(w, word) && !visited[i] {
				path = append(path, w)
				visited[i] = true
			}
		}
	}

	return ans
}

func isDiffer1(s, w string) bool {
	var count int
	for i := 0; i < len(w); i++ {
		if s[i] != w[i] {
			count++
		}
	}
	if count == 1 {
		return true
	}
	return false
}
