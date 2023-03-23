package lc

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func ClimbingStairs(n int) int {
	//if n <= 2 {
	//	return n
	//}
	//dp := make([]int, n+1)
	//for i := 0; i < n+1; i++ {
	//	dp[i] = 1
	//}
	//for i := 2; i <= n; i++ {
	//	dp[i] = dp[i-1] + dp[i-2]
	//}
	//return dp[n]
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func HouseRobber(nums []int) int {
	//n := len(nums)
	//dp := make([]int, n)
	//dp[0] = nums[0]
	//for i := 1; i < n; i++ {
	//	if i-2 < 0 {
	//		dp[i] = max(dp[i-1], nums[i])
	//	} else {
	//		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	//	}
	//
	//}
	//return max(dp[n-2], dp[n-1])

	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	var pre1, pre2, cur int
	for i := 0; i < n; i++ {
		cur = max(pre2+nums[i], pre1)
		pre2 = pre1
		pre1 = cur
	}
	return cur
}

func NumberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	count, ans := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]+nums[i-2] == 2*nums[i-1] {
			count++
			ans += count
		} else {
			count = 0
		}
	}

	return ans
}

func MinPathSum(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}

		}
	}

	return dp[m-1][n-1]
}
func ZeroOneMatrix1(nums [][]int) [][]int {
	m, n := len(nums), len(nums[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums[i][j] == 0 {
				dp[i][j] = 0
			} else {
				if j > 0 {
					dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
				}
				if i > 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
				}
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums[i][j] != 0 {
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
			}
		}
	}

	return dp
}

func ZeroOneMatrix2(nums [][]byte) [][]int {
	directions := []int{-1, 0, 1, 0, -1}
	m, n := len(nums), len(nums[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var queue [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				dp[i][j] = math.MaxInt32
			}
		}
	}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x, y := pos[0]+directions[i], pos[1]+directions[i+1]
			if x >= 0 && x < m && y >= 0 && y < n {
				if nums[x][y] == 1 && dp[x][y] == math.MaxInt32 {
					dp[x][y] = min(dp[x][y], dp[pos[0]][pos[1]]+1)
					queue = append(queue, []int{x, y})
				}
			}
		}
	}
	return dp
}

func MaximalSquare(nums [][]byte) int {
	m, n := len(nums), len(nums[0])
	var maxSide int
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
			maxSide = max(dp[i][j], maxSide)
		}
	}
	return maxSide * maxSide
}

func MaximalSquareBrute(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return 0
	}
	maxSide := 0
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				currMaxSide := min(m-1, n-1)
				for k := 1; k < currMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for t := 0; t < k; t++ {
						if matrix[i+k][j+t] == '0' || matrix[i+t][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}

			}
		}
	}
	return maxSide * maxSide
}

func NumSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= n; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func NumDecoding(s string) int {
	return numDecodingDFS(s, 0)
}

func numDecodingDFS(s string, pos int) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	pre := s[0] - '0'
	if pre == '0' {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}
	for i := 2; i <= n; i++ {
		curr := s[i-1] - '0'
		if (pre == 0 || pre > 2) && curr == 0 {
			return 0
		}
		if (pre == 1) || pre == 2 && curr < 7 {
			if curr > 0 {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
		pre = curr
	}
	return dp[n]
}

func WordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]
}
func LongestIncreasingSubsequence(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	var maxLen int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

func LongestCommonSequence(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)
	C := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		C[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				C[i][j] = C[i-1][j-1] + 1

			} else if C[i-1][j] >= C[i][j-1] {
				C[i][j] = C[i-1][j]
			} else {
				C[i][j] = C[i][j-1]
			}
		}
	}
	return C[l1][l2]
}

func CanPartition(nums []int) bool {
	var sum int
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 || n < 2 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j < target+1; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][target]

}

func KnapsackTwoDimension(values, weights []int, W int) int {
	vLen, wLen := len(values), len(weights)
	if vLen <= 0 || wLen <= 0 {
		return 0
	}
	dp := make([][]int, vLen+1)
	for i := 0; i < vLen+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= vLen; i++ {
		v, w := values[i-1], weights[i-1]
		for j := 1; j <= W; j++ {
			if j >= w {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w]+v)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[vLen][W]
}

func KnapsackOneDimension(values, weights []int, W int) int {
	wLen := len(weights)
	vLen := len(values)
	dp := make([]int, wLen+1)
	for i := 1; i <= vLen; i++ {
		w, v := weights[i-1], values[i-1]
		for j := W; j >= w; j-- {
			dp[j] = max(dp[j], dp[j-w]+v)
		}
	}
	return dp[wLen]
}

func KnapsackNPTwoDimension(values, weights []int, W int) int {
	n := len(values)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++ {
		w, v := weights[i-1], values[i-1]
		for j := 1; j <= W; j++ {
			if j >= w {
				dp[i][j] = max(dp[i-1][j], dp[i][j-w]+v)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][W]
}

func CoinsChange(coins []int, amount int) int {
	n := len(coins)
	if n <= 0 {
		return -1
	}
	dp := make([]int, amount+2)
	dp[0] = 0
	for i := 1; i < amount+2; i++ {
		dp[i] = amount + 2
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+2 {
		return -1
	}
	return dp[amount]
}

func BuyAndSellSock(sock []int) int {
	n := len(sock)
	minVal := sock[0]
	var maxVal int
	for i := 1; i < n; i++ {
		if sock[i] < minVal {
			minVal = sock[i]
		}
		maxVal = max(maxVal, sock[i]-minVal)
	}
	return maxVal
}

func DiffWaysToCompute(input string) []int {
	var ways []int
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c == '-' || c == '+' || c == '*' {
			left := DiffWaysToCompute(input[:i])
			right := DiffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch c {
					case '-':
						ways = append(ways, l-r)
						break
					case '+':
						ways = append(ways, l+r)
						break
					case '*':
						ways = append(ways, l*r)
						break

					}
				}
			}

		}

	}
	return ways
}

func EratosthenesPrime(n int) int {
	if n < 2 {
		return 0
	}
	composite := make([]bool, n)
	count := n - 2
	for i := 2; i < n; i++ {
		if !composite[i] {
			for j := 2 * i; j < n; j += i {
				if !composite[j] {
					composite[j] = true
					count--
				}
			}
		}
	}
	return count
}
func CountsPrimes(n int) int {
	if n < 2 {
		return 0
	}
	var count int
	var flag bool
	for i := 2; i < n; i++ {
		flag = false
		sq := int(math.Sqrt(float64(i)))
		if sq == 1 {
			count++
			continue
		}
		for j := 2; j <= sq; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}
		if !flag {
			count++
		}
	}
	return count
}

func BaseSeven(n int) string {
	n1 := n
	if n < 0 {
		n1 = -n
	}
	nu := 7
	re := n1 % nu
	qu := n1 / nu
	var ans string
	for ; qu != 0; qu = qu / nu {
		ans = strconv.Itoa(re) + ans
		re = qu % nu
	}

	ans = strconv.Itoa(re) + ans
	if n < 0 {
		ans = "-" + ans
	}
	return ans

}

func TrailingZeroes(n int) int {
	//var count int
	//for i := 1; i <= n; i++ {
	//	if i%5 == 0 {
	//		count++
	//	}
	//}
	//return count
	if n == 0 {
		return 0
	}
	return n/5 + TrailingZeroes(n/5)
}
func PowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	ans := n / 3
	y := ans % 3
	for ; ans >= 3; ans = ans / 3 {
		if y != 0 {
			return false
		}
		y = ans % 3
	}
	if ans == 1 && y == 0 {
		return true
	}
	return false
}

func IsIsomorphic(s string, t string) bool {
	ss, ts := make([]int, 256), make([]int, 256)
	for i := 0; i < len(s); i++ {
		if ss[s[i]] != ts[t[i]] {
			return false
		}
		ss[s[i]], ts[t[i]] = i+1, i+1
	}
	return true
}

func CountSubstring(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		count += extendSubstring(s, i, i)
		count += extendSubstring(s, i, i+1)
	}
	return count
}

func extendSubstring(s string, l, r int) int {
	var count int
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		count++
	}
	return count
}

func BasicCalculator(s string) int {

	return 0
}

func priority(s, t byte) bool {
	if (s == '+' || s == '-') && (t == '*' || t == '\\') {
		return false
	}
	return true
}

func SubString(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	i, j, k := 0, 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			j++
			i++
		} else {
			k++
			i = k
			j = 0
		}
	}
	if j == len(needle) {
		return k
	}
	return -1
}

func IndexKMP(haystack string, needle string) int {
	next := calNext(needle)
	n, m, k := len(haystack), len(needle), -1
	if m <= 0 {
		return -1
	}
	for i := 0; i < n; i++ {
		for k > -1 && needle[k+1] != haystack[i] {
			k = next[k]
		}
		if needle[k+1] == haystack[i] {
			k++
		}
		if k == m-1 {
			return i - m + 1
		}
	}

	return -1
}

func calNext(needle string) []int {
	n, p := len(needle), -1
	next := make([]int, n)
	for i := 0; i < n; i++ {
		next[i] = -1
	}
	for i := 1; i < n; i++ {
		for p > -1 && needle[p+1] != needle[i] {
			p = next[p]
		}
		if needle[p+1] == needle[i] {
			p++
		}
		next[i] = p
	}
	fmt.Println(next)
	return next
}
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	var c int
	for i := 1; i < n; i++ {
		if contains(s[c:i], s[i]) {
			dp[i] = dp[i-1]
		} else {
			count := 1
			j := i
			for ; j >= 1; j-- {
				if !contains(s[j:i+1], s[j-1]) {
					count++
				} else {
					break
				}
			}
			if count > dp[i-1] {
				c = j
				dp[i] = count
			} else {
				dp[i] = dp[i-1]
			}

		}
	}
	return dp[n-1]
}

func contains(s string, t byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == t {
			return true
		}
	}
	return false
}

func MaxDepth(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return 1 + max(MaxDepth(tree.Left), MaxDepth(tree.Right))
}

func IsBalance(root *TreeNode) bool {
	return isBalHelper(root) != -1
}

func isBalHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := isBalHelper(root.Left), isBalHelper(root.Right)
	if left == -1 || right == -1 || abs(left, right) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func DiameterOfBinaryTree(root *TreeNode) int {
	var d int
	diaHelper(root, &d)
	return d
}

func diaHelper(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}
	left, right := diaHelper(root.Left, d), diaHelper(root.Right, d)
	*d = max(left+right, *d)
	return max(left, right) + 1
}

func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := rootSum(root, targetSum)
	ans += PathSum(root.Left, targetSum-root.Val)
	ans += PathSum(root.Right, targetSum-root.Val)
	return ans
}

func rootSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		ans++
	}
	ans += rootSum(root.Left, targetSum-root.Val)
	ans += rootSum(root.Right, targetSum-root.Val)
	return
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelper(root.Left, root.Right)
}
func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	var ans []float64
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		var sum int
		t := len(queue)
		for i := 0; i < t; i++ {
			r := queue[0]
			queue = queue[1:]
			sum += r.Val
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		ans = append(ans, float64(sum)/float64(t))
	}
	return ans
}

func BuildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return buildTreeHelper(m, preorder, 0, len(preorder)-1, 0)
}

func buildTreeHelper(m map[int]int, preorder []int, s0, e0, s1 int) *TreeNode {
	if s0 > e0 {
		return nil
	}
	mid := preorder[s1]
	index := m[mid]
	leftLen := index - s0 - 1
	node := new(TreeNode)
	node.Val = mid
	node.Left = buildTreeHelper(m, preorder, s0, index-1, s1+1)
	node.Right = buildTreeHelper(m, preorder, index+1, e0, s1+2+leftLen)
	return node
}

func PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := new([]TreeNode)
	var ans []int
	*stack = append(*stack, *root)
	for len(*stack) > 0 {
		r := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		ans = append(ans, r.Val)
		if r.Right != nil {
			*stack = append(*stack, *r.Right)
		}
		if r.Left != nil {
			*stack = append(*stack, *r.Left)
		}
	}
	return ans
}

func RecoverTree(root *TreeNode) {
	var stack []*TreeNode
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val

}
