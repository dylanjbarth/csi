package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

type bloomFilter interface {
	add(item string)

	// `false` means the item is definitely not in the set
	// `true` means the item might be in the set
	maybeContains(item string) bool

	// Number of bytes used in any underlying storage
	memoryUsage() int
}

// Use case:
// optimize for storing _very roughly_ half of the dictionary
// $ wc -l /usr/share/dict/words
// 235886 /usr/share/dict/words so ballpark N ~100k items.

/*
	false positive rate: (1-e^(-kn/m))^k
	where n is the size of our set
	and k is the number of hash functions.
	and m is the number of bits in our bit array.

	Optimization, assuming n ~ 100k:

k	          m	      n	       kb used
0.003465736	500	    100000	 0.06
0.006931472	1000	  100000	 0.13
0.013862944	2000	  100000	 0.25
0.027725887	4000	  100000	 0.50
0.055451774	8000	  100000	 1.00
0.110903549	16000	  100000	 2.00
0.221807098	32000	  100000	 4.00
0.443614196	64000	  100000	 8.00
0.887228391	128000	100000	 16.00
1.774456782	256000	100000	 32.00
3.548913564	512000	100000	 64.00
7.097827129	1024000	100000	 128.00
14.19565426	2048000	100000	 256.00
*/

type awesomeBloomFilter struct {
	data []uint64
	m    uint64
	k    uint64
}

func newAwesomeBloomFilter(m, k uint64) *awesomeBloomFilter {
	if m%64 != 0 {
		panic(fmt.Sprintf("m must be a factor of 64 to make the math easy (lazy!!). Received %d", m))
	}
	return &awesomeBloomFilter{
		data: make([]uint64, m/64),
		m:    m,
		k:    k,
	}
}

func (b *awesomeBloomFilter) add(item string) {
	idx, bitToSet := b.getBits(item)
	b.data[idx] |= bitToSet
	// fmt.Printf("%s => %d mod %d => %d\n", item, hashed, b.m, bitToSet)
	// fmt.Printf("data[%d] flipping bit at index %d\n", idx, innerIdx)
	// fmt.Printf("%b\n", b.data)
}

func (b *awesomeBloomFilter) maybeContains(item string) bool {
	idx, bitToSet := b.getBits(item)
	return (b.data[idx] & bitToSet) > 0
}

func (b *awesomeBloomFilter) memoryUsage() int {
	return binary.Size(b.data)
}

func (b *awesomeBloomFilter) getBits(item string) (uint64, uint64) {
	hasher := fnv.New64()
	hasher.Write([]byte(item))
	hashed := hasher.Sum64()
	bitToSet := hashed % b.m
	idx := bitToSet / 64
	innerIdx := bitToSet - (idx * 64)
	bitSet := 1 << innerIdx
	return idx, uint64(bitSet)
}
