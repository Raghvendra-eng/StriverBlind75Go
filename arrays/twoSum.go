package arrays

import "sort"

/*

Problem Statement: Given an array of integers arr[] and an integer target.

1st variant: Return YES if there exist two numbers such that their sum is equal to the target. Otherwise, return NO.

2nd variant: Return indices of the two numbers such that their sum is equal to the target. Otherwise, we will return {-1, -1}.

Note: You are not allowed to use the same element twice. Example: If the target is equal to 6 and num[1] = 3, then nums[1] + nums[1] = target is not a solution.

Examples:

Example 1:
Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 14
Result: YES (for 1st variant)
       [1, 3] (for 2nd variant)
Explanation: arr[1] + arr[3] = 14. So, the answer is “YES” for the first variant and [1, 3] for 2nd variant.

Example 2:
Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 15
Result: NO (for 1st variant)
	[-1, -1] (for 2nd variant)
Explanation: There exist no such two numbers whose sum is equal to the target.

*/

// If array is not sorted, Use map.
// Time Complexity O(n)
// Space Complexity O(n)
func checkTwoSum(arr []int, target int) (bool, []int) {
	hm := make(map[int]int)
	ans := []int{-1, -1}
	for index, num := range arr {
		if val, ok := hm[target-num]; ok {
			ans[0] = val
			ans[1] = index
			return true, ans
		}
		hm[arr[index]] = index
	}
	return false, ans
}


// If array is sorted, use two pointers
func checkTwoSumSortedArray(arr []int, target int) (bool, []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	ans := []int{-1, -1}
	for i, j := 0, len(arr) - 1; i < j; i++ {
		if (target == (arr[i] + arr[j])) {
			ans[0] = i
			ans[1] = j
			return true, ans
		} else if (target > (arr[i] + arr[j])) {
			i++
		} else {
			j--
		}
	}
	return false, ans
}
