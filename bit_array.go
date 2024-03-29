package beef

const (
	BIT      uint64 = 1
	INT_SIZE uint64 = 64
)

type BitArray struct {
	Size uint64
	data []uint64
}

// New takes in a size and outputs a pointer to a new BitArray.
func NewBitArray(size uint64) *BitArray {
	ba := new(BitArray)
	ba.Size = size
	ba.data = make([]uint64, (size/INT_SIZE)+1)
	return ba
}

// Sets the bit array bit at the input position to 1.
func (b *BitArray) Set(position uint64) {
	integer := position / b.Size
	bit_position := position % INT_SIZE
	b.data[integer] = b.data[integer] | (BIT << bit_position)
}

// Returns the bit at the input position.
func (b *BitArray) Get(position uint64) uint64 {
	integer := position / b.Size
	bit_position := position % INT_SIZE
	return (b.data[integer] & (BIT << bit_position)) >> bit_position
}
