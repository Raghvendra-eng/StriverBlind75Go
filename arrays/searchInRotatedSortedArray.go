package arrays

/**
Problem Statement: Given an integer array arr of size N,
sorted in ascending order (with distinct values) and a target value k.
Now the array is rotated at some pivot point unknown to you.
Find the index at which k is present and if k is not present return -1.
**/

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
