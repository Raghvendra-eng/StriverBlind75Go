package arrays

// Problem Statement: Given an array that contains both negative and positive integers, find the maximum product subarray.

func maxProduct(nums []int) int {
	ans := nums[0]
	smallest := nums[0]
	greatest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			smallest, greatest = greatest, smallest
		}
		smallest = smallest * nums[i]
		greatest = greatest * nums[i]
		if smallest > nums[i] {
			smallest = nums[i]
		}
		if greatest < nums[i] {
			greatest = nums[i]
		}
		ans = max(greatest, ans)
	}
	return ans
}
