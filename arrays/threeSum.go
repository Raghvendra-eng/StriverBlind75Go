package arrays

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/**
Problem Statement: Given an array of N integers,
your task is to find unique triplets that add up to give a sum of zero.
In short, you need to return an array of all the unique triplets
[arr[a], arr[b], arr[c]] such that i!=j, j!=k, k!=i, and their sum is equal to zero.
**/

func threeSum(nums []int) [][]int {
	ans := make(map[string]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[i]+nums[j]+nums[k] == 0 {
				entry := fmt.Sprint(nums[i]) + ":" + fmt.Sprint(nums[j]) + ":" + fmt.Sprint(nums[k])
				ans[entry] = true
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return getAnswer(ans)
}

func getAnswer(ans map[string]bool) [][]int {
	var answer [][]int
	for key, _ := range ans {
		stringValue := strings.Split(key, ":")
		fmt.Println(stringValue)
		var value []int
		for _, num := range stringValue {
			a, _ := strconv.Atoi(num)
			value = append(value, a)
		}
		answer = append(answer, value)
	}
	return answer
}
