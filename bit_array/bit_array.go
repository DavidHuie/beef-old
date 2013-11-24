package bit_array

const bit uint64 = 1
const bits uint64 = 64

type BitArray struct {
	size uint64
	data []uint64
}

// New takes in a size and outputs a pointer to a new BitArray.
func New(size uint64) *BitArray {
	ba := new(BitArray)
	ba.size = size
	ba.data = make([]uint64, (size/bits)+1)
	return ba
}

// Sets the bit array bit at the input position to 1.
func (b *BitArray) Set(position uint64) {
	integer := position / b.size
	bit_position := position % bits
	b.data[integer] = b.data[integer] | (bit << bit_position)
}

// Returns the bit at the input position.
func (b *BitArray) Get(position uint64) uint64 {
	integer := position / b.size
	bit := position % bits
	return (b.data[integer] & (uint64(1) << bit)) >> bit
}
