package binary

/**
Reverse bits of a given 32 bits signed integer.
**/
func reverseBits(n int) int {
    ans := 0
	for i := 31; i >= 0; i-- {
		ans = (ans<<1) | ((n>>(31-i)) & 1)
	}
	return ans 
}