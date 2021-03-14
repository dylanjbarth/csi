package i2b

// Return stack with binary representation of decimal number n
func i2b(n int) []int {
	b := make([]int, 0)
	val := n / 2
	rem := n % 2
	b = append(b, rem)
	for val > 0 {
		rem = val % 2
		val /= 2
		b = append(b, rem)
	}
	return b
}
