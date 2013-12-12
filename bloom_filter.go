package beef

type BloomFilter struct {
	bit_array *BitArray
	hashes    uint
}

func NewBloomFilter(size uint64, hashes uint) *BloomFilter {
	bf := new(BloomFilter)
	bf.bit_array = NewBitArray(size)
	bf.hashes = hashes
	return bf
}

func (b *BloomFilter) Insert(value string) {
	for _, value := range b.hash_values(value) {
		b.bit_array.Set(value)
	}
}

func (b *BloomFilter) Check(value string) bool {
	for _, value := range b.hash_values(value) {
		check := b.bit_array.Get(value)
		if check != BIT {
			return false
		}
	}
	return true
}

func (b *BloomFilter) hash_values(value string) []uint64 {
	values := make([]uint64, 0)
	for i := uint(0); i < b.hashes; i++ {
		values = append(values, Hash(value, i)%b.bit_array.Size)
	}
	return values
}
