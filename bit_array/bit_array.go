package bit_array

const bits uint64 = 64

type BitArray struct {
	size uint64
	data []uint64
}

func expected_size(size uint64) uint64 {
	// This is an estimate.
	return (size / bits) + 1
}

// New takes in a size and outputs a pointer to a new BitArray.
func New(size uint64) *BitArray {
	ba := new(BitArray)
	ba.size = size
	ba.data = make([]uint64, expected_size(size))
	return ba
}

func (b *BitArray) Set(position uint64) {
	integer := position / b.size
	bit := position % bits
	b.data[integer] = b.data[integer] & (uint64(1) << bit)
}

func (b *BitArray) Get(position uint64) uint64 {
	integer := position / b.size
	bit := position % bits
	return (b.data[integer] & (uint64(1) << bit)) >> bit
}
