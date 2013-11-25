package bloom_filter

import (
	"github.com/DavidHuie/beef/bit_array"
	"github.com/DavidHuie/beef/string_hash"
)

type BloomFilter struct {
	bit_array *bit_array.BitArray
	hashes    uint
}

func New(size uint64, hashes uint) *BloomFilter {
	bf := new(BloomFilter)
	bf.bit_array = bit_array.New(size)
	bf.hashes = hashes
	return bf
}

func (b *BloomFilter) Insert(value string) {
	for _, value := range hash_values(value, b.hashes) {
		b.bit_array.Set(value)
	}
}

func (b *BloomFilter) Check(value string) bool {
	for _, value := range hash_values(value, b.hashes) {
		check := b.bit_array.Get(value)
		if check != bit_array.BIT {
			return false
		}
	}
	return true
}

func hash_values(value string, hashes uint) []uint64 {
	values := make([]uint64, 0)
	for i := uint(0); i < hashes; i++ {
		values = append(values, string_hash.Hash(value, i))
	}
	return values
}
