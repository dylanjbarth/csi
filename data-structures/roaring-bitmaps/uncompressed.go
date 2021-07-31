package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	// 2^32 / 64 => max capacity of the bitmap
	return &uncompressedBitmap{data: make([]uint64, 67_108_864)}
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	idx := x / 64
	bitToCheck := x % 64
	return b.data[idx]&(1<<bitToCheck) > 0
}

func (b *uncompressedBitmap) Set(x uint32) {
	// Identify the bitmap index (where the decimal value is the index of the bit to flip)
	idx := x / 64
	bitToSet := x % 64
	// Set the bit by taking the bitwise OR of the existing value and the bit that we need to set.
	b.data[idx] = b.data[idx] | (1 << bitToSet)
}

// TODO Union & intersection would be more annoying if we didn't know the bitmaps were the same length...
func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	var data = make([]uint64, len(b.data))
	for i := 0; i < len(b.data); i++ {
		data[i] = b.data[i] | other.data[i]
	}
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	var data = make([]uint64, len(b.data))
	for i := 0; i < len(b.data); i++ {
		data[i] = b.data[i] & other.data[i]
	}
	return &uncompressedBitmap{
		data: data,
	}
}
