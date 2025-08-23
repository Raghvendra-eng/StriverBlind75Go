package arrays

/**
You are given an integer array height of length n.
There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
**/

func maxArea(height []int) int {
	ans := 0
	for i, j := 0, len(height)-1; i < j; {
		water := min(height[i], height[j]) * (j - i)
		ans = max(ans, water)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
