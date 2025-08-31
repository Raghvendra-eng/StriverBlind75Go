package dynamicprogamming

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	rememberAnswer := []int{}
	for i := 0; i <= amount; i++ {
		rememberAnswer = append(rememberAnswer, math.MaxInt)
	}
	rememberAnswer[0] = 0
	for _, coin := range coins {
		if coin < amount {
			rememberAnswer[coin] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		if rememberAnswer[i] == 1 {
			continue
		}
		for _, coin := range coins {
			if i-coin >= 0 && rememberAnswer[i-coin] != math.MaxInt {
				rememberAnswer[i] = min(rememberAnswer[i], rememberAnswer[i-coin]+1)
			}
		}
	}
	if rememberAnswer[amount] == math.MaxInt {
		return -1
	}
	return rememberAnswer[amount]
}
