package bitmap

import (
	// "fmt"
	"fmt"
	"math"
)

// Word Aligned Hybrid encoding, 1 signal - 63 following bits.
func compress(b *uncompressedBitmap) []uint64 {
	var out []uint64
	n_runs := 0
	offset := uint64(0)
	chunksize := 63
	n_iterations := int(math.Ceil(float64(len(b.data)) * 64 / float64(chunksize)))
	b.PrettyPrint()
	for i := 0; i < n_iterations; i++ {
		chunk := getNext63Bits(b, offset)
		fmt.Printf("i: %d Chunk: %064b\n", i, chunk)
		if chunk == 0 {
			n_runs += 1
			offset += uint64(chunksize)
			continue
		} else if n_runs > 0 {
			// flush n_runs before processing this chunk.
			fmt.Printf("i: %d Flushing %d 0s => %064b\n", i, n_runs, uint64(0))
			out = append(out, uint64(n_runs))
			n_runs = 0
		}
		// flip signal bit to indicate literal and then write the rest of the 63 bits here
		toWrite := chunk | 1<<63
		fmt.Printf("Adding: %064b\n", toWrite)
		out = append(out, toWrite)
		offset += uint64(chunksize)
	}
	return out
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
	end := start + 63 // End index isn't inclusive
	var val uint64
	// Each element in b.data is 64 bits, so we need to find the starting place and overflow to the next element if needed.
	b1_idx := start / 64
	b1_inner_idx := start % 64
	b2_idx := end / 64
	b2_inner_idx := end % 64
	fmt.Printf("Bitoffset, start, end: %d %d %d \n", bitoffset, start, end)
	fmt.Printf("Block index: [b%d:%d-b%d:%d] \n", b1_idx, b1_inner_idx, b2_idx, b2_inner_idx)
	if b1_idx == b2_idx { // in this rare situation, we can just extract the element and clean off the starting or ending bit.
		fmt.Printf("Single element approach:\n")
		fmt.Printf("Element: %064b\n", b.data[b1_idx])
		if b1_inner_idx == 0 { // clear end bit
			val = b.data[b1_idx] >> 1
		} else { // clear starting bit
			val = b.data[b1_idx] &^ (1 << 63)
		}
		fmt.Printf("Val: %064b\n", val)
	} else { // in the more common situation we have to join bits from neighboring indexes, clearing bits from both
		fmt.Printf("Two element approach:\n")
		fmt.Printf("First: %064b\n", b.data[b1_idx])
		first := b.data[b1_idx] << b1_inner_idx
		fmt.Printf("First shifted: %064b\n", first)
		second := uint64(0)
		if len(b.data)-1 > int(b2_idx) {
			fmt.Printf("Second: %064b\n", second)
			second = b.data[b2_idx] >> b2_inner_idx
			fmt.Printf("Second shifted: %064b\n", second)
		} else {
			fmt.Printf("Second: %064b\n", second)
			fmt.Printf("Second shifted: %064b\n", second)
		}
		// create a complete uint64 from the two indexes.
		fmt.Printf("Val: %064b\n", val)
		val = first | second
		// Finally, shift right to only return 63 bits.
		val = val >> 1
		fmt.Printf("Val: %064b\n", val)
	}
	return val
}
