package binary

/**
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), 
ans[i] is the number of 1's in the binary representation of i.
**/

func countBits(n int) []int {
	ans := []int{}
	ans = append(ans, 0)
	if n == 0 {
		return ans
	}
	ans = append(ans, 1)
	nextPowerOf2 := 2
	for i := 2; i <= n; i++ {
		if i == nextPowerOf2 {
			ans = append(ans, 1)
			nextPowerOf2 *= 2
		} else {
			base := nextPowerOf2/2
			ans = append(ans, 1 + ans[i-base])
		}
	}
	return ans
}