package beef

const (
	// The size of an underlying bit array
	// is this number times the expected size.
	bit_array_factor = 12
	// The number of hash functions use in
	// new bloom filters.
	num_hashes = 3
)

type Manager struct {
	bit_arrays map[string]*BloomFilter
}

func NewManager() *Manager {
	bs := new(Manager)
	bs.bit_arrays = make(map[string]*BloomFilter)
	return bs
}

func (b *Manager) CreateBF(name string, size uint64) {
	b.bit_arrays[name] = NewBloomFilter(size*bit_array_factor, num_hashes)
}

func (b *Manager) DeleteBF(name string) {
	delete(b.bit_arrays, name)
}

func (b *Manager) ExistsBF(name string) bool {
	_, ok := b.bit_arrays[name]
	return ok
}

func (b *Manager) InsertBF(name string, value string) {
	b.bit_arrays[name].Insert(value)
}

func (b *Manager) CheckBF(name string, value string) bool {
	return b.bit_arrays[name].Check(value)
}
