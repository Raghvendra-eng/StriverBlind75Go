package arrays

import (
	"sort"
)

/**
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example:

Example 1:
Input: nums = [1, 2, 3, 1]
Output: true.
Explanation: 1 appeared two times in the input array.

Example 2:
Input: nums = [1, 2, 3, 4]
Output: false
Explanation: input array does not contain any duplicate number.
*/

// Time complexity: O(n)
// Space complexity: O(n)
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}

// Time complexity: O(n log n)
// Space complexity: O(1)
func containsDuplicateO1Space(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
