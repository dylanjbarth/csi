package bitmap

import (
	"math"
)

// Word Aligned Hybrid encoding, 1 signal - 63 following bits.
func compress(b *uncompressedBitmap) []uint64 {
	var out []uint64
	n_runs := 0
	offset := uint64(0)
	chunksize := uint64(63)
	n_iterations := int(math.Ceil(float64(len(b.data)) * 64 / float64(chunksize)))
	b.PrettyPrint()
	for i := 0; i < n_iterations; i++ {
		chunk := getNext63Bits(b, offset)
		// fmt.Printf("i: %d Chunk: %064b\n", i, chunk)
		if chunk == 0 {
			n_runs += 1
			offset += chunksize
			continue
		} else if n_runs > 0 {
			// flush n_runs before processing this chunk.
			// fmt.Printf("i: %d Flushing %d 0s => %064b\n", i, n_runs, uint64(0))
			out = append(out, uint64(n_runs))
			n_runs = 0
		}
		// flip signal bit to indicate literal and then write the rest of the 63 bits here
		toWrite := chunk | 1<<63
		// fmt.Printf("Adding: %064b\n", toWrite)
		out = append(out, toWrite)
		offset += chunksize
	}
	return out
}

func decompress(compressed []uint64) *uncompressedBitmap {
	var data []uint64
	var tmpBits uint64
	var nBits uint64
	for i := 0; i < len(compressed); i++ {
		curr := compressed[i]
		signalBit := curr & (1 << 63)
		// fmt.Printf("curr %064b\n", curr)
		// fmt.Printf("signal bit %064b\n", signalBit)
		if signalBit > 0 { // next 63 bits are literal
			if nBits == 0 {
				tmpBits = curr << 1
				nBits = 63
				// fmt.Printf("tmpBits %064b\n", tmpBits)
			} else {
				// flush tmpBits
				// clear signal bit
				curr = curr &^ (1 << 63)
				// fmt.Printf("curr %064b\n", curr)
				emptySlots := 64 - nBits
				toAdd := tmpBits | (curr >> (63 - emptySlots))
				data = append(data, toAdd)
				// fmt.Printf("toAdd %064b\n", toAdd)
				// now that we flushed, store the other half
				tmpBits = curr << emptySlots
				nBits = 63 - emptySlots
				// fmt.Printf("tmpBits %064b\n", tmpBits)
			}
		} else { // this is a count of runs of 64 0s
			// flush the runs of 0s
			for j := 0; j < int(curr); j++ {
				if nBits == 0 {
					tmpBits = 0
					nBits = 63
					// fmt.Printf("tmpBits %064b\n", tmpBits)
				} else {
					// flush tmpBits
					emptySlots := 64 - nBits
					data = append(data, tmpBits)
					// now that we flushed, store the other half
					tmpBits = 0
					nBits = 63 - emptySlots
					// fmt.Printf("tmpBits %064b\n", tmpBits)
				}
			}
		}
	}
	// Final flush (don't flush zeroes?)
	if nBits > 0 && tmpBits > 0 {
		data = append(data, tmpBits)
	}
	return &uncompressedBitmap{
		data: data,
	}
}

// Returns next 63 bits from bit offset as a uint64 (where the most significant bit will be 0 and can be ignored)
func getNext63Bits(b *uncompressedBitmap, bitoffset uint64) uint64 {
	start := bitoffset
	end := start + 63
	var val uint64
	// Each element in b.data is 64 bits, so we need to find the starting place and overflow to the next element if needed.
	b1_idx := start / 64
	b1_inner_idx := start % 64
	b2_idx := end / 64
	// b2_inner_idx := end % 64
	// fmt.Printf("Bitoffset, start, end: %d %d %d \n", bitoffset, start, end)
	// fmt.Printf("Block index: [b%d:%d-b%d:%d] \n", b1_idx, b1_inner_idx, b2_idx, b2_inner_idx)
	if b1_idx == b2_idx { // in this rare situation, we can just extract the element and clean off the starting or ending bit.
		// fmt.Printf("Single element approach:\n")
		// fmt.Printf("Element: %064b\n", b.data[b1_idx])
		if b1_inner_idx == 0 { // clear end bit
			val = b.data[b1_idx] &^ (1)
		} else { // clear starting bit
			val = b.data[b1_idx] << 1
		}
	} else { // in the more common situation we have to join bits from neighboring indexes, clearing bits from both
		// fmt.Printf("Two element approach:\n")
		// fmt.Printf("First: %064b\n", b.data[b1_idx])
		first := b.data[b1_idx] << b1_inner_idx
		// fmt.Printf("First shifted: %064b\n", first)
		second := uint64(0)
		if len(b.data)-1 >= int(b2_idx) {
			// fmt.Printf("Second: %064b\n", b.data[b2_idx])
			second = b.data[b2_idx] >> (64 - b1_inner_idx)
			// fmt.Printf("Second shifted: %064b\n", second)
		} else {
			// fmt.Printf("Second: %064b\n", second)
			// fmt.Printf("Second shifted: %064b\n", second)
		}
		// create a complete uint64 from the two indexes, 0ing out end bit.
		val = (first | second) &^ (1)
	}
	// As a final step, shift right to ignore the most significant bit.
	// fmt.Printf("Val: %064b\n", val)
	return val >> 1
}
