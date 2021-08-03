package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	// Original approach was to pre-allocate, but this is really slow to compress if it's not needed...
	// 2^32 / 64 => max capacity of the bitmap, 8MB
	// return &uncompressedBitmap{data: make([]uint64, 67_108_864)}
	// Try empty strategy so that compress is faster
	return &uncompressedBitmap{data: []uint64{}}
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	idx := x / 64
	bitToCheck := x % 64
	if int(idx) > len(b.data)-1 {
		return false
	}
	return b.data[idx]&(1<<bitToCheck) > 0
}

func (b *uncompressedBitmap) Set(x uint32) {
	// Identify the bitmap index (where the decimal value is the index of the bit to flip)
	idx := x / 64
	bitToSet := x % 64
	// Extend length of slice if needed.
	diff := int(idx) - len(b.data) + 1
	if diff > 0 {
		b.data = append(b.data, make([]uint64, diff)...)
	}
	// Set the bit by taking the bitwise OR of the existing value and the bit that we need to set.
	b.data[idx] = b.data[idx] | (1 << bitToSet)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	minL := min(len(b.data), len(other.data))
	var data = make([]uint64, minL)
	for i := 0; i < minL; i++ {
		data[i] = b.data[i] | other.data[i]
	}
	// fill in the rest
	largest := b
	if len(b.data) < len(other.data) {
		largest = other
	}
	data = append(data, (largest.data[minL:])...)
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	minL := min(len(b.data), len(other.data))
	var data = make([]uint64, minL)
	for i := 0; i < minL; i++ {
		data[i] = b.data[i] & other.data[i]
	}
	return &uncompressedBitmap{
		data: data,
	}
}
