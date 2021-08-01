package bitmap

import "fmt"

// Word Aligned Hybrid encoding, 1 signal - 63 following bits.
func compress(b *uncompressedBitmap) []uint64 {
	// n_runs := 0
	// var out []uint64
	chunk := getNext63Bits(b, 0)
	// iterations = len(b.data) / 64
	for i := 0; i < len(b.data); i++ {

	}
	return []uint64{chunk}
}

func decompress(compressed []uint64) *uncompressedBitmap {
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

// Returns next 63 bits from bit offset as a uint64
func getNext63Bits(b *uncompressedBitmap, bitoffset uint64) uint64 {
	start := bitoffset
	end := start + 63
	var val uint64
	// Each element in b.data is 64 bits, so we need to find the starting place and overflow to the next element if needed.
	s_idx := start / 64
	s_rem := start % 64
	e_idx := end / 64
	e_rem := end % 64
	fmt.Printf("Bitoffset, start, end: %d %d %d \n", bitoffset, start, end)
	fmt.Printf("Start index, end index: %d %d \n", s_idx, e_idx)
	fmt.Printf("Start remainder, end remainder: %d %d \n", s_rem, e_rem)
	if s_idx == e_idx { // in this rare situation, we can just extract the element and clean off the starting or ending bit.
		fmt.Printf("Single element approach:\n")
		fmt.Printf("Element: %b\n", b.data[s_idx])
		if s_rem == 0 { // clear end bit
			val = b.data[s_idx] >> 1
		} else { // clear starting bit
			val = b.data[s_idx] &^ (1 << 63)
		}
		fmt.Printf("Val: %b\n", val)
	} else { // in the more common situation we have to join bits from two adjacent indexes, clearing bits from both
		fmt.Printf("Two element approach:\n")
		first := b.data[s_idx] << s_rem
		second := uint64(0)
		if len(b.data)-1 <= int(e_idx) {
			second = b.data[e_idx] >> e_rem
		}
		val = first | second
		fmt.Printf("First: %b\n", b.data[s_idx])
		fmt.Printf("First shifted: %b\n", first)
		fmt.Printf("Second: %b\n", b.data[e_idx])
		fmt.Printf("Second shifted: %b\n", second)
		fmt.Printf("Val: %b\n", val)
	}
	return val
}
