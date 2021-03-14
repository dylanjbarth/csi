package i2b

import (
	"log"
	"sort"
)

// Converts a decimal integer to binary representation
func i2b(n int) []int {
	b := make([]int, 0)
	val := n / 2
	rem := n % 2
	b = append(b, rem)
	for val > 0 {
		log.Printf("Value is %d", rem)
		log.Printf("Remainder is %d", rem)
		rem = val % 2
		val /= 2
		b = append(b, rem)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	return b
}
