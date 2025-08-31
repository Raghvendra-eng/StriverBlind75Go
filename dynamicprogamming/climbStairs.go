package dynamicprogamming

/**
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
**/

// Time Complexity O(n)
// Space Complexity O(n)

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	ans := []int{}
	ans = append(ans, 1)
	ans = append(ans, 2)
	for i := 2; i < n; i++ {
		ans = append(ans, ans[i-1]+ans[i-2])
	}
	return ans[n-1]
}

// nth fabonaki no

// Time Complexity O(n)
// Space Complexity O(1)

func climbStairs1(n int) int {
	if n == 0 {
		return 0
	}
	curr, prev := 2, 1
	for i := 3; i <= n; i++ {
		temp := curr
		curr += prev
		prev = temp
	}
	return curr
}

// Time Complexity O(log n)
// Space Complexity O(1)
func climbStairs2(n int) int {
	if n == 0 {
		return 0
	}
	ans := &[][]int{{1, 1}, {1, 0}}
	ans = power(ans, n)
	return (*ans)[0][0]
}

func power(mat *[][]int, n int) *[][]int {
	if n == 0 {
		return &[][]int{{1, 0}, {0, 1}}
	}
	ans := power(mat, n/2)
	ans = multiply(ans, ans)
	if n & 1 == 1 {
		ans = multiply(ans, mat)
	}
	return ans
}

func multiply(mat1, mat2 *[][]int) *[][]int {
	ans := [][]int{{0, 0}, {0, 0}}
	ans[0][0] = (*mat1)[0][0] * (*mat2)[0][0] + (*mat1)[0][1] * (*mat2)[1][0]
	ans[0][1] = (*mat1)[0][0] * (*mat2)[0][1] + (*mat1)[0][1] * (*mat2)[1][1]
	ans[1][0] = (*mat1)[1][0] * (*mat2)[0][0] + (*mat1)[1][1] * (*mat2)[1][0]
	ans[1][1] = (*mat1)[1][0] * (*mat2)[0][1] + (*mat1)[1][1] * (*mat2)[1][1]
	return &ans
}
