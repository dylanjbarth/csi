// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	bd := BitDiff32(c1, c2)
	fmt.Printf("\nBitDiff %d", bd)
}

// BitDiff32 returns the number of bits that are different between two byte arrays.
func BitDiff32(b1 [32]byte, b2 [32]byte) int {
	fmt.Println("First byte array")
	for _, n := range b1 {
		fmt.Printf("%08b", n)
	}
	fmt.Println("")
	fmt.Println("Second byte array")
	for _, n := range b2 {
		fmt.Printf("%08b", n)
	}
	count := 0
	for i := range b1 {
		for b := 0; b < 7; b++ {
			// compare the rightmost bit in each byte and count the differences.
			b1b := b1[i] >> b & 1
			b2b := b2[i] >> b & 1
			if b1b != b2b {
				count++
			}
		}
	}
	return count
}
