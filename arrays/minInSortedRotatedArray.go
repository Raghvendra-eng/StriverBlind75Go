package arrays

/**
Problem Statement: Given an integer array arr of size N,
sorted in ascending order (with distinct values).
Now the array is rotated between 1 to N times which is unknown. Find the minimum element in the array.
**/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]

}
