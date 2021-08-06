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
	for i := 0; i < n_iterations; i++ {
		chunk := getNext63Bits(b, offset)
		if chunk == 0 {
			n_runs += 1
			offset += chunksize
			continue
		} else if n_runs > 0 {
			// flush n_runs before processing this chunk.
			out = append(out, uint64(n_runs))
			n_runs = 0
		}
		// flip signal bit to indicate literal and then write the rest of the 63 bits here
		toWrite := chunk | 1<<63
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
		if signalBit > 0 { // next 63 bits are literal
			if nBits == 0 {
				tmpBits = curr << 1
				nBits = 63
			} else {
				// flush tmpBits
				// clear signal bit
				curr = curr &^ (1 << 63)
				emptySlots := 64 - nBits
				toAdd := tmpBits | (curr >> (63 - emptySlots)) // add one to account for the signal bit
				data = append(data, toAdd)
				// now that we flushed, store the other half
				tmpBits = curr << (emptySlots + 1)
				nBits = 63 - emptySlots
			}
		} else { // this is a count of runs of 64 0s
			// flush the runs of 0s
			for j := 0; j < int(curr); j++ {
				if nBits == 0 {
					tmpBits = 0
					nBits = 63
				} else {
					// flush tmpBits
					emptySlots := 64 - nBits
					data = append(data, tmpBits)
					// now that we flushed, store the other half
					tmpBits = 0
					nBits = 63 - emptySlots
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
	if b1_idx == b2_idx { // in this rare situation, we can just extract the element and clean off the starting or ending bit.
		if b1_inner_idx == 0 { // clear end bit
			val = b.data[b1_idx] &^ (1)
		} else { // clear starting bit
			val = b.data[b1_idx] << 1
		}
	} else { // in the more common situation we have to join bits from neighboring indexes, clearing bits from both
		first := b.data[b1_idx] << b1_inner_idx
		second := uint64(0)
		if len(b.data)-1 >= int(b2_idx) {
			second = b.data[b2_idx] >> (64 - b1_inner_idx)
		} else {
		}
		// create a complete uint64 from the two indexes, 0ing out end bit.
		val = (first | second) &^ (1)
	}
	// As a final step, shift right to ignore the most significant bit.
	return val >> 1
}
