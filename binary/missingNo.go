package binary

/**
Given an array nums containing n distinct numbers in the range [0, n], 
return the only number in the range that is missing from the array.
**/

func missingNumber(nums []int) int {
	ans := 0
    for i, n := range nums {
		ans = ans ^ i ^ n
	}
	return ans ^ len(nums)
}