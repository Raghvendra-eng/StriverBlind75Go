package binary

/**
Given two integers a and b, return the sum of the two integers without using the operators + and -
**/

func getSum(a int, b int) int {
    for b > 0 {
        carry := (a & b) << 1
        a = a ^ b
        b = carry
    }
    return a
}