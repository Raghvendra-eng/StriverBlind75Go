package binary
/**
Given a positive integer n, write a function that returns the number 
of set bits in its binary representation (also known as the Hamming weight).
 **/

 func hammingWeight(n int) int {
	ans := 0
	for n > 0 {
		ans++
		n = n & (n-1) // convert least significant set bit to 0
	} 
	return ans
 }