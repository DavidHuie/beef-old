package beef

import "errors"

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

func (b *Manager) CreateBF(name string, size uint64) error {
	if _, ok := b.bit_arrays[name]; ok {
		return errors.New("Bloom filter already exists")
	}
	b.bit_arrays[name] = NewBloomFilter(size*bit_array_factor, num_hashes)
	return nil
}

func (b *Manager) DeleteBF(name string) error {
	if _, ok := b.bit_arrays[name]; !ok {
		return errors.New("Bloom filter does not exist")
	}
	delete(b.bit_arrays, name)
	return nil
}

func (b *Manager) ExistsBF(name string) bool {
	_, ok := b.bit_arrays[name]
	return ok
}

func (b *Manager) InsertBF(name string, value string) error {
	if _, ok := b.bit_arrays[name]; !ok {
		return errors.New("Bloom filter does not exist")
	}
	b.bit_arrays[name].Insert(value)
	return nil
}

func (b *Manager) CheckBF(name string, value string) (bool, error) {
	if _, ok := b.bit_arrays[name]; !ok {
		return false, errors.New("Bloom filter does not exist")
	}
	return b.bit_arrays[name].Check(value), nil
}
