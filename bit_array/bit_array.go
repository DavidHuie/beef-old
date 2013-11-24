package bit_array

const BIT uint64 = 1
const BITS uint64 = 64

type BitArray struct {
	size uint64
	data []uint64
}

// New takes in a size and outputs a pointer to a new BitArray.
func New(size uint64) *BitArray {
	ba := new(BitArray)
	ba.size = size
	ba.data = make([]uint64, (size/BITS)+1)
	return ba
}

// Sets the bit array bit at the input position to 1.
func (b *BitArray) Set(position uint64) {
	integer := position / b.size
	bit_position := position % BITS
	b.data[integer] = b.data[integer] | (BIT << bit_position)
}

// Returns the bit at the input position.
func (b *BitArray) Get(position uint64) uint64 {
	integer := position / b.size
	bit_position := position % BITS
	return (b.data[integer] & (BIT << bit_position)) >> bit_position
}
